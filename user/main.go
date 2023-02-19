package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	userpb "handson/user/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s server) Get(context.Context, *userpb.GetRequest) (*userpb.GetResponse, error) {
	return &userpb.GetResponse{
		Name: "taro",
	}, nil
}

func main() {
	closer := serve()
	defer closer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func serve() (closer func()) {
	log.Println("User Service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	go func() {
		log.Printf("Server started at %v", lis.Addr().String())
		reflection.Register(s)
		err = s.Serve(lis)
		if err != nil {
			log.Println("ERROR:", err.Error())
		}
	}()

	closer = func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		log.Println("stopping gRPC server...")
		s.Stop()
	}
	return closer
}
