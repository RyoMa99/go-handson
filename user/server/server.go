package server

import (
	"context"
	"handson/user/db"
	userpb "handson/user/proto"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init(userDB db.UserDB) (closer func()) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	return serve(lis, userDB)
}

type server struct {
	userpb.UnimplementedUserServiceServer
	userDB db.UserDB
}

func serve(lis net.Listener, userDB db.UserDB) (closer func()) {
	log.Println("User Service")

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{userDB: userDB})

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
	c, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	uid := primitive.NewObjectID()

	err := s.userDB.UpsertOne(c, &db.User{
		Id:   uid,
		Name: in.Name,
		Age:  in.Age,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &userpb.CreateUserResponse{
		Id: uid.Hex(),
	}, nil
}
