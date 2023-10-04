package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	context "context"

	itemProto "./_proto/api/"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	itemProto.RegisterItemServiceServer(server, NewMyServer())

	reflection.Register(server)

	go func() {
		log.Printf("grpc server start %v", port)
		server.Serve(listener)
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	log.Println("stop grpc server")
	server.GracefulStop()

}

type myServer struct {
	itemProto.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *itemProto.ItemRequest) (*itemProto.ItemResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &itemProto.HelloResponse{
		Message: fmt.Sprintf("Hello, item request id %s!", req.GetId()),
	}, nil
}
