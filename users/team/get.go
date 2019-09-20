package team

import (
	"context"
	"log"

	"obscura-users-backend/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbUsers "obscura-users-backend/proto"
)

func (TeamService) GetTeam(ctx context.Context, req *pbUsers.GetTeamRequest) (*pbUsers.GetTeamResponse, error) {
	teamID, err := primitive.ObjectIDFromHex(req.GetTeamID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	filter := bson.D{{"_id", teamID}}

	var team db.Team
	if err := db.TeamCollection.FindOne(context.TODO(), filter).Decode(&team); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	log.Println("team is")
	log.Println(team)

	requests := []*pbUsers.Requests{}
	for _, r := range team.Requests {
		log.Println(r)
		requests = append(requests, &pbUsers.Requests{ID: r.ID.Hex(), Name: r.Name, ImageNumber: r.ImageNumber})
	}
	log.Println("requests are")
	log.Println(requests)
	log.Println(team.Requests)
	return &pbUsers.GetTeamResponse{
		ID:          team.ID.Hex(),
		Name:        team.Name,
		Level:       int32(team.Level),
		CreatorID:   team.CreatorID.Hex(),
		Requests:    requests,
		ImageNumber: team.ImageNumber,
	}, nil
}
