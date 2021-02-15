package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	helloworldpb "github.com/akhripko/grpc-gateway/api/helloworld"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

// func NewServer() *server {
// 	return &server{}
// }

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	var token string
	headers, ok := metadata.FromIncomingContext(ctx)
	if ok {
		h := headers["my-proto-header"]
		if len(h) == 0 {
			return nil, errors.New("my header was not provided")
		}
		token = h[0]
		log.Println("my-proto-header: ", token)
	}
	header := metadata.New(map[string]string{"my-srv-proto-header": "srv-" + token})
	grpc.SendHeader(ctx, header)

	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
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
	helloworldpb.RegisterGreeterServer(s, &server{})
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
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
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

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")
	w.Header().Set("Content-Type", "application/json")

	buf, merr := marshaler.Marshal(Error{
		Code:    100500,
		Message: err.Error(),
	})
	if merr != nil {
		log.Println("failed to marshal error message: ", merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			log.Println("failed to write response: ", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("failed to write response: %v", err)
	}
}
