package main

import (
	context "context"

	// authpb "github.com/RaymondAkachi/VAULT-API/proto"
	authpb "github.com/RaymondAkachi/VAULT-API/proto"
)

type AuthServer struct {
	authpb.UnimplementedAUthServiceServer
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
	err := CreateUser(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthResponse{
		Message: "Registration successful",
		Token: "fake-jwt-token", // Eeplace with real jwt later
	}, nil


}	

