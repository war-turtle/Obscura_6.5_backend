package team

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"obscura-users-backend/db"

	pbUsers "obscura-users-backend/proto"
)

// GetAllTeams function to get all teams info
func (TeamService) GetAllTeams(ctx context.Context, req *pbUsers.GetAllTeamsRequest) (*pbUsers.GetAllTeamsResponse, error) {
	log.Println(req.GetPage())
	findOptions := options.Find().SetLimit(10).SetSort(bson.D{{"level", -1}}).SetSkip(int64(10 *  (req.GetPage() - 1)))
	cur, err := db.TeamCollection.Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	defer cur.Close(ctx)

	var teams []*pbUsers.TeamInfo
	for cur.Next(ctx) {
		var elem map[string]interface{}
		if err := cur.Decode(&elem); err != nil {
			log.Println(err)
		}
		teams = append(teams, &pbUsers.TeamInfo{ID: elem["_id"].(primitive.ObjectID).Hex(), Name: elem["name"].(string), Level: elem["level"].(int32), ImageNumber: elem["imagenumber"].(int32) })
	}

	totalTeams, err := db.TeamCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	return &pbUsers.GetAllTeamsResponse{Teams: teams, Total: int32(totalTeams)}, nil
}