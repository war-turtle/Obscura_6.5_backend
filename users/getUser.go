package main

import (
	"context"
	"log"

	"obscura-users-backend/db"
	"obscura-users-backend/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	pbUsers "obscura-users-backend/proto"
)

func (*server) GetUser(ctx context.Context, req *pbUsers.GetUserRequest) (*pbUsers.JwtResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.GetUserID())
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}
	var user db.User
	filter := bson.D{{"_id", userID}};
	if err := db.UserCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated user")
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