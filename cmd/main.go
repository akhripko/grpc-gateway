package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/akhripko/grpc-gateway/api/echo"
)

type server struct {
	pb.UnimplementedEchoServiceServer
}

// func NewServer() *server {
// 	return &server{}
// }

func (s *server) PostEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var token string
	headers, ok := metadata.FromIncomingContext(ctx)
	if ok {
		h := headers["my-proto-header"]
		if len(h) == 0 {
			st, _ := status.New(codes.PermissionDenied, "my-proto-header was not provided").
				WithDetails(&pb.Error{
					Code:    403,
					Message: "My-Header was not provided",
				})
			return &pb.EchoResponse{}, st.Err()
		}
		token = h[0]
		log.Println("my-proto-header: ", token)
	}
	header := metadata.New(map[string]string{"my-srv-proto-header": "srv-" + token})
	grpc.SendHeader(ctx, header)

	return &pb.EchoResponse{
		Name:  in.Name,
		Data1: in.Data1,
		Data2: in.Data2,
		EmId:  in.EmId,
	}, nil
}

func (s *server) GetEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Name:  in.Name,
		Data1: in.Data1,
		Data2: in.Data2,
		EmId:  in.EmId,
	}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", "0.0.0.0:8080") // nolint
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
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

	// Register Greeter
	err = pb.RegisterEchoServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

func IncomingHeaderMatcher(key string) (string, bool) {
	if key == "My-Header" {
		return "My-Proto-Header", true
	}
	return runtime.DefaultHeaderMatcher(key)
}

func OutgoingHeaderMatcher(key string) (string, bool) {
	if key == "my-srv-proto-header" {
		return "My-Srv-Header", true
	}
	return "", false
}

type Error struct {
	Code    int
	Message string
}

func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// return Internal when Marshal failed
	const fallback = `{"code": 13, "message": "failed to marshal error message"}`

	var errMsg *pb.Error
	st, ok := status.FromError(err)
	if ok {
		errMsg = StatusToError(st)
	}
	if errMsg == nil {
		errMsg = &pb.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(errMsg.Code))
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
