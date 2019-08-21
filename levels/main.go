package main

import (
	"log"
	"net"

	pbLevels "obscura-levels-backend/proto"

	"google.golang.org/grpc"
	"obscura-levels-backend/db"
	configuration "obscura-levels-backend/config"
)

type server struct{}

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
	pbLevels.RegisterLevelServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
