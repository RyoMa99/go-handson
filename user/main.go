package main

import (
	"context"
	"log"
	"net"

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
	log.Println("User Service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	log.Printf("Server started at %v", lis.Addr().String())
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Println("ERROR:", err.Error())
	}
}
