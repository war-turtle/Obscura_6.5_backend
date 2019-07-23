package team

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"obscura-users-backend/db"
	"obscura-users-backend/jwt"
	
	configuration "obscura-users-backend/config"
	pbUsers "obscura-users-backend/proto"
)

// TeamService struct for team methods
type TeamService struct{}

var config = configuration.GetConfig()

// CreateTeam function to create a new team
func (TeamService) CreateTeam(ctx context.Context, req *pbUsers.CreateTeamRequest) (*pbUsers.JwtResponse, error) {
	log.Println(req)

	var team db.Team
	filter := bson.D{{"name", req.GetName()}}
	if err := db.TeamCollection.FindOne(context.TODO(), filter).Decode(&team); err != nil {
		if err == mongo.ErrNoDocuments {
			creatorID, err := primitive.ObjectIDFromHex(req.GetCreatorID())
			if err != nil {
				return nil, status.Error(codes.Unknown, "internal server error")
			}

			team = db.Team {
				Name: req.GetName(),
				CreatorID: creatorID,
				Requests: []db.Requests{},
			}
	
			res, err := db.TeamCollection.InsertOne(context.TODO(), team)
			if err != nil {
				return nil, status.Error(codes.Unknown, "internal server error")
			}
			team.ID = res.InsertedID.(primitive.ObjectID)

			filter := bson.D{{"_id", creatorID}}
			update := bson.D{
				{
					"$set", bson.D{
						{
							"teamid", team.ID,
						},
					},
				},
			}

			var user db.User
			opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
			if err := db.UserCollection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&user); err != nil {
				return nil, status.Error(codes.Unknown, "internal server error");
			}

			tokenSecret, err := jwt.CreateNewUserJwt(user, config["jwtSecret"])
			if err != nil {
				return nil, status.Error(codes.Unknown, "internal server error")
			}
	
			return &pbUsers.JwtResponse{Jwt: tokenSecret}, nil
		}

		return nil, status.Error(codes.Unknown, "internal server error")
	}
	
	return nil, status.Error(codes.Unknown, "team already exists")
}

