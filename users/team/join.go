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

var channels = make(map[string](chan int))

func (TeamService) JoinTeam(ctx context.Context, req *pbUsers.JoinTeamRequest) (*pbUsers.Empty, error) {
	userID, err := primitive.ObjectIDFromHex(req.GetUserID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	teamID, err := primitive.ObjectIDFromHex(req.GetTeamID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	var user db.User
	filter := bson.D{{"_id", userID}}
	if err := db.UserCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	var team db.Team
	filter = bson.D{{"_id", teamID}}
	if err := db.TeamCollection.FindOne(context.TODO(), filter).Decode(&team); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	for _, r := range team.Requests {
		if r.ID == userID {
			return nil, status.Error(codes.Unknown, "request already sent")
		}
	}

	update := bson.D{{"$push", bson.D{{"requests", bson.D{{"id", user.ID}, {"name", user.Username}, {"imagenumber", user.ImageNumber}}}}}}
	if _, err := db.TeamCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	return &pbUsers.Empty{}, nil
}
