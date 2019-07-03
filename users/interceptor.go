package main

import (
	"log"
	"strings"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"obscura-users-backend/jwt"
)

func valid(authorization string) bool {
	if len(authorization) < 1 {
		return false
	}

	token := strings.Split(authorization, " ")
	if(token[0] == "google"){
		return true
	} else if (token[0] == "jwt") {
		return jwt.ValidateJwt(token[1], config["jwtSecret"])
	} else {
		return false
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"][0]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return m, err
}