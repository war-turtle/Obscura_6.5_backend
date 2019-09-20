package main

import (
	"log"
	"net"

	pbUsers "obscura-users-backend/proto"

	configuration "obscura-users-backend/config"
	"obscura-users-backend/db"
	"obscura-users-backend/team"

	"google.golang.org/grpc"
)

type server struct {
	team.TeamService
}

var config = configuration.GetConfig()

func init() {
	db.Connect(config["mongoURI"])
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
