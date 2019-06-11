package main;

import (
	"context"
	"net/http"
	"io/ioutil"
	"encoding/json"

	pbUsers "obscura-users-backend/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*server) SignIn(ctx context.Context, req *pbUsers.SignInRequest) (*pbUsers.SignInResponse, error) {
	response, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token="+req.IdToken);
	if err != nil {
		return nil, status.Error(codes.Unknown, "can't reach google server")
	}

	var data map[string]interface{}

	res, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(res, &data)
	
	if val, ok := data["error"]; ok {
		return nil, status.Error(codes.Unauthenticated, val.(string))
	}
	
	return &pbUsers.SignInResponse{Name: data["name"].(string), Email: data["email"].(string)}, nil
}