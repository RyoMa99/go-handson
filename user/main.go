package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("User Service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	s := grpc.NewServer()
	// s.RegisterService()

	log.Printf("Server started at %v", lis.Addr().String())
	err = s.Serve(lis)
	if err != nil {
		log.Println("ERROR:", err.Error())
	}
}
