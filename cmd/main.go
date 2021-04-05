package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/akhripko/grpc-gateway/api/echo"
	"github.com/akhripko/grpc-gateway/pkg"
)

type server struct {
	pb.UnimplementedEchoServiceServer
}

// func NewServer() *server {
// 	return &server{}
// }

func (s *server) PostEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	cookie := metadata.New(map[string]string{"set-cookie": "srvcookie:cookie_value"})
	grpc.SetHeader(ctx, cookie)

	var token string
	headers, ok := metadata.FromIncomingContext(ctx)
	if ok {
		h := headers["my-header"]
		if len(h) == 0 {
			st, _ := status.New(codes.PermissionDenied, "my-proto-header was not provided").
				WithDetails(&pb.Error{
					Code:    403,
					Message: "My-Header was not provided",
				})
			return &pb.EchoResponse{}, st.Err()
		}
		token = h[0]
		log.Println("my-header: ", token)
	}
	header := metadata.New(map[string]string{"my-srv-header": "srv-" + token})
	grpc.SetHeader(ctx, header)

	return &pb.EchoResponse{
		Name:    in.Name,
		Data1:   in.Data1,
		Data2:   in.Data2,
		EmId:    in.EmId,
		BoolVal: in.BoolVal,
	}, nil
}

func (s *server) GetEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Name:    in.Name,
		Data1:   in.Data1,
		Data2:   in.Data2,
		EmId:    in.EmId,
		BoolVal: in.BoolVal,
	}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", "0.0.0.0:8080") // nolint
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	var s = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(addTraceIDUnaryInterceptor)),
	)
	pb.RegisterEchoServiceServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(IncomingHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(OutgoingHeaderMatcher),
		runtime.WithErrorHandler(ErrorHandler))

	withLogging := pkg.WithLoggingMiddleware(gwmux)

	// Register Greeter
	err = pb.RegisterEchoServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: withLogging,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

type void struct{}

var member = void{}

var outgoingHeaders = map[string]void{
	"set-cookie":    member,
	"x-trace-id":    member,
	"my-srv-header": member,
}

func OutgoingHeaderMatcher(key string) (string, bool) {
	if key == "grpcgateway-content-type" {
		return "", false
	}
	lower := strings.ToLower(key)
	if _, ok := outgoingHeaders[lower]; ok {
		return lower, true
	}
	return runtime.DefaultHeaderMatcher(lower)
}

var incomingHeaders = map[string]void{
	"x-trace-id": member,
	"my-header":  member,
}

func IncomingHeaderMatcher(key string) (string, bool) {
	lower := strings.ToLower(key)
	if _, ok := incomingHeaders[lower]; ok {
		return lower, true
	}
	return runtime.DefaultHeaderMatcher(key)
}

type Error struct {
	Code    int
	Message string
}

func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// return Internal when Marshal failed
	const fallback = `{"code": 13, "message": "failed to marshal error message"}`

	st, _ := status.FromError(err)
	errMsg := StatusToError(st)

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}
	if md.HeaderMD != nil {
		trID := md.HeaderMD.Get("x-trace-id")
		if len(trID) > 0 {
			errMsg.TraceID = trID[0]
		}
	}
	handleForwardResponseServerMetadata(w, md)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	buf, merr := marshaler.Marshal(errMsg)
	if merr != nil {
		log.Println("failed to marshal error message: ", merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			log.Println("failed to write response: ", err)
		}
		return
	}

	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("failed to write response: %v", err)
	}
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.HeaderMD {
		if h, ok := OutgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}
}

func StatusToError(st *status.Status) *pb.Error {
	if st == nil {
		return nil
	}
	details := st.Details()
	if len(details) != 0 {
		if res, ok := details[0].(*pb.Error); ok {
			return res
		}
	}
	return &pb.Error{Code: 500, Message: st.Message()}
}

func addTraceIDUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// get id from metadata
	id, ok := ReadMetadataValue(ctx, "x-trace-id")
	if !ok {
		return handler(ctx, req)
	}
	// set atomic trace id header
	grpc.SetHeader(ctx, metadata.New(map[string]string{
		"x-trace-id": id,
	}))

	return handler(ctx, req)
}

func ReadMetadataValue(ctx context.Context, key string) (string, bool) {
	// get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	// get value from metadata
	v, ok := md[key]
	if !ok || len(v) == 0 {
		return "", true
	}
	return v[0], true
}
