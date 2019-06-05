package main

import (
	"context"
	"log"
	"net"

	pbLevels "obscura-proto/levels"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetLevels(ctx context.Context, req *pbLevels.LevelRequest) (*pbLevels.LevelResponse, error) {
	log.Printf("Received: %v", req.Id)
	return &pbLevels.LevelResponse{Id: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Can't listen %v", err)
	}

	s := grpc.NewServer()

	pbLevels.RegisterLevelServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
