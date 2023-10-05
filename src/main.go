package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	context "context"

	itemProto "grpc_bff_backend/src/_proto/api"

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
	itemProto.UnimplementedItemServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *itemProto.ItemRequest) (*itemProto.ItemResponse, error) {
	return &itemProto.ItemResponse{
		Id:       0,
		ItemNo:   0,
		Name:     "",
		Quantity: 0,
		Unit:     "",
		Price:    0,
		Remark:   "",
	}, nil
}
