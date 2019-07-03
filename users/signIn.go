package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"obscura-users-backend/db"
	"obscura-users-backend/jwt"

	pbUsers "obscura-users-backend/proto"
)

func (*server) SignIn(ctx context.Context, req *pbUsers.SignInRequest) (*pbUsers.JwtResponse, error) {
	response, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + req.IdToken)
	if err != nil {
		return nil, status.Error(codes.Unknown, "can't reach google server")
	}

	var data map[string]interface{}

	res, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(res, &data)

	if val, ok := data["error"]; ok {
		return nil, status.Error(codes.Unauthenticated, val.(string))
	}

	email := data["email"].(string);

	var user db.User
	filter := bson.D{{"email", email}}
	if err := db.Collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		if(err == mongo.ErrNoDocuments){
			username := data["name"].(string);

			user = db.User{
				Username: username,
				Email: email,
				Phone: "",
				College: "",
				Onboard: false,
			}

			if _, err := db.Collection.InsertOne(context.TODO(), &user); err != nil{
				return nil, status.Error(codes.Unknown, "internal server error")
			}

		} else {
			log.Fatalln(err)
		}
	}

	tokenString, err := jwt.CreateNewUserJwt(user, config["jwtSecret"])
	if err != nil {
		return nil, status.Error(codes.Unknown, "internal server error")
	}

	log.Println(user)

	return &pbUsers.JwtResponse{Jwt: tokenString}, nil
}
