package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pbPlayers "obscura-proto/players"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetPlayers(ctx context.Context, req *pbPlayers.PlayerRequest) (*pbPlayers.PlayerResponse, error) {
	fmt.Println(req.Id)
	return &pbPlayers.PlayerResponse{Id: 2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("can't listen %v", err)
	}

	// creds, err := credentials.NewServerTLSFromFile("/etc/ssl/certs/selfsigned.crt", "/etc/ssl/private/selfsigned.key")
	// if err != nil {
	// 	log.Fatalf("failed to create credentials: %v", err)
	// }

	s := grpc.NewServer()
	pbPlayers.RegisterPlayerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
