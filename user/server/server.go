package server

import (
	"context"
	userpb "handson/user/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() (closer func()) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	return serve(lis)
}

type server struct {
	userpb.UnimplementedUserServiceServer
}

func serve(lis net.Listener) (closer func()) {
	log.Println("User Service")

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	go func() {
		log.Printf("Server started at %v", lis.Addr().String())
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Println("ERROR:", err.Error())
		}
	}()

	closer = func() {
		log.Println("stopping gRPC server...")
		s.GracefulStop()
	}
	return closer
}

func (s server) Get(context.Context, *userpb.GetRequest) (*userpb.GetResponse, error) {
	return &userpb.GetResponse{
		Name: "taro",
	}, nil
}

func (s server) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return &userpb.CreateUserResponse{
		Id: in.Id,
	}, nil
}
