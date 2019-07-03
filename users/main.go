package main

import (
	"log"
	"net"

	pbUsers "obscura-users-backend/proto"

	"google.golang.org/grpc"
	"obscura-users-backend/db"
)

type server struct{}

var config = GetConfig()

func init() {
	go db.Connect(config["mongoURI"])
}

func main() {
	lis, err := net.Listen("tcp", ":"+config["port"])
	if err != nil {
		log.Fatalf("can't listen %v", err)
	} else {
		log.Println("Server is listen on port " + config["port"])
	}

	log.Println("server started")
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	pbUsers.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
