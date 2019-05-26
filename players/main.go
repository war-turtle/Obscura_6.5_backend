package main

import (
	"context"
	"log"
	"net"

	pbPlayers "github.com/war-turtle/Obscura_6.5_backend/proto/players"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetPlayers(ctx context.Context, req *pbPlayers.PlayerRequest) (*pbPlayers.PlayerResponse, error) {
	return &pbPlayers.PlayerResponse{Id: 2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("can't listen %v", err)
	}

	s := grpc.NewServer()
	pbPlayers.RegisterPlayerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
