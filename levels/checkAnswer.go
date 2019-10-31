package main

import (
	"context"

	"obscura-levels-backend/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbLevels "obscura-levels-backend/proto"
)

func (server) CheckAnswer(ctx context.Context, req *pbLevels.AnswerRequest) (*pbLevels.AnswerResponse, error) {
	var level db.Level
	filter := bson.D{{"number", req.GetNumber()}}
	if err := db.LevelCollection.FindOne(context.TODO(), filter).Decode(&level); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	for _, ans := range level.Ans {
		if req.GetAnswer() == ans {
			teamID, err := primitive.ObjectIDFromHex(req.GetTeamID())
			if err != nil {
				return nil, status.Error(codes.Unknown, "internal server error")
			}

			filter = bson.D{{"_id", teamID}}

			var team db.Team
			if err := db.TeamCollection.FindOne(context.TODO(), filter).Decode(&team); err != nil {
				return nil, status.Error(codes.Unknown, "internal server error")
			}
			if int32(team.Level) == req.GetNumber() {
				update := bson.D{{"$inc", bson.D{{"level", 1}}}}
				if err := db.TeamCollection.FindOneAndUpdate(context.TODO(), filter, update).Err(); err != nil {
					return nil, status.Error(codes.Unknown, "internal server error")
				}
			}
			return &pbLevels.AnswerResponse{Valid: true}, nil
		}
	}
	return nil, status.Error(codes.PermissionDenied, "wrong answer")
}
