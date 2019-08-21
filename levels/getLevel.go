package main

import (
	"log"
	"context"

	pbLevels "obscura-levels-backend/proto"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server) GetLevel(ctx context.Context, req *pbLevels.LevelRequest) (*pbLevels.LevelResponse, error) {
	log.Println("wow")
	return nil, status.Error(codes.Unknown, "internal server error")
}