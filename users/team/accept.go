package team

import (
	"context"

	"obscura-users-backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbUsers "obscura-users-backend/proto"
)

func (TeamService) AcceptUser(ctx context.Context,req *pbUsers.JoinTeamRequest) (*pbUsers.Empty, error) {
	userID, err := primitive.ObjectIDFromHex(req.GetUserID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "invalid user")
	}
	teamID, err := primitive.ObjectIDFromHex(req.GetTeamID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	var user db.User
	filter := bson.D{{ "_id", userID }}
	if err := db.UserCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	if user.TeamID.IsZero() {
		update := bson.D{{ "$set", bson.D{{ "teamid", teamID }} }}
		if _, err := db.UserCollection.UpdateOne(context.TODO(), filter, update); err != nil {
			return nil, status.Error(codes.Unknown, "internal server error")
		}
		filter = bson.D{{ "_id", teamID }}
		update = bson.D{{ "$pull", bson.D{{ "requests", bson.D{{ "id", userID }} }} }}
		if err := db.TeamCollection.FindOneAndUpdate(context.TODO(), filter, update).Err(); err != nil {
			return nil, status.Error(codes.Unknown, "internal server error")
		}
		return &pbUsers.Empty{}, nil
	}

	filter = bson.D{{ "_id", teamID }}
	update := bson.D{{ "$pull", bson.D{{ "requests", bson.D{{ "id", userID }} }} }}
	if err := db.TeamCollection.FindOneAndUpdate(context.TODO(), filter, update).Err(); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	return nil, status.Error(codes.Unknown, "player has already joined another team")
}