package main

import (
	"context"
	"log"
	
	"obscura-users-backend/db"
	"obscura-users-backend/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pbUsers "obscura-users-backend/proto"
)

func (*server) OnBoard(ctx context.Context, req *pbUsers.OnBoardRequest) (*pbUsers.JwtResponse, error) {
	filter := bson.D{{"email", req.Email}};
	update := bson.D{
		{
			"$set", bson.D{
				{ "username", req.Username },
				{ "college", req.College },
				{ "phone", req.Phone },
				{ "onboard", true },
				{ "imagenumber", req.ImageNumber},
			},
		},
	}
	
	var user db.User
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := db.UserCollection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Unauthenticated, "wrong email address")
		}
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	log.Println(user)

	tokenSecret, err := jwt.CreateNewUserJwt(user, config["jwtSecret"])
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	return &pbUsers.JwtResponse{Jwt: tokenSecret}, nil
}