package main

import (
	"context"

	"obscura-levels-backend/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbLevels "obscura-levels-backend/proto"
)

func (server) GetLevel(ctx context.Context, req *pbLevels.LevelRequest) (*pbLevels.LevelResponse, error) {
	teamID, err := primitive.ObjectIDFromHex(req.GetTeamID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	filter := bson.D{{"_id", teamID}}

	var team db.Team
	if err := db.TeamCollection.FindOne(context.TODO(), filter).Decode(&team); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	if int32(team.Level) < req.GetId() {
		return nil, status.Error(codes.OutOfRange, "level unreachable")
	}

	var level db.Level
	filter = bson.D{{"number", req.GetId()}}
	if err := db.LevelCollection.FindOne(context.TODO(), filter).Decode(&level); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "no such level found")
		}
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	return &pbLevels.LevelResponse{Number: level.Number, Name: level.Name, Urlalias: level.UrlAlias, Html: level.Html, Js: level.Js, Final: level.Final}, nil
}
