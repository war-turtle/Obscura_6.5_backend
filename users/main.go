package main

import (
	"log"
	"net"

	pbUsers "obscura-users-backend/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("can't listen %v", err)
	}

	s := grpc.NewServer()
	pbUsers.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
