package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"1

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"obscura-levels-backend/db"
	pbLevels "obscura-levels-backend/proto"
)

func (server) GetAllLevel(ctx context.Context, req *pbLevels.AllLevelRequest) (*pbLevels.AllLevelResponse, error) {
	filter := bson.D{{}}
	total, err := db.LevelCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	return &pbLevels.AllLevelResponse{Total: int32(total)}, nil
}
