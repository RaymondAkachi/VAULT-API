package main

import (
	context "context"
	"fmt"
	"time"

	db "github.com/RaymondAkachi/VAULT-API/auth-service/internal/database"
	utils "github.com/RaymondAkachi/VAULT-API/auth-service/utils"
	authpb "github.com/RaymondAkachi/VAULT-API/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
    authpb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
    user, err := CreateUser(ctx, req.Username, req.Email, req.Password)
    if err != nil {
        return nil, err
    }

    // Create JWT token 

    token, err := utils.GenerateToken(req.Email, 24*time.Hour)
    if err != nil {
        return nil, err
    }

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token")
	}
    
    created_token, rt_error := queries.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	if rt_error != nil {
		return nil, rt_error
	}

    return &authpb.AuthResponse{
        Message: "Registration successful",
        Token:  token,
		RefreshToken: created_token.Token, // Replace with real JWT later
    }, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
    user, err := AuthenticateUser(ctx, req.Email, req.Password)
    if err != nil {
        return nil, err
    }
    
    // Create JWT Token

    token, err := utils.GenerateToken(req.Email, 24*time.Hour)
    if err != nil{
        return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
    }

    refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token")
	}
    
    created_token, err := queries.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save refresh token")
	}

    return &authpb.AuthResponse{
        Message: "Login successful",
        Token:   token,
		RefreshToken: created_token.Token, // Replace with real JWT later
    }, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *authpb.TokenRequest) (*authpb.ValidateTokenResponse, error) {
	claims, err := utils.ValidateToken(req.Token)
	if err != nil {
		return &authpb.ValidateTokenResponse{Valid: false}, nil
	}

	return &authpb.ValidateTokenResponse{
		Valid: true,
		Email: claims.Email,
	}, nil
}

func (s *AuthServer) RefreshToken(ctx context.Context, req *authpb.RefreshRequest) (*authpb.AuthResponse, error) {
	rtoken, err := queries.GetRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("refresh token invalid or expired")
	}

	if rtoken.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("refresh token expired")
	}

	user, err := queries.GetUserByID(ctx, rtoken.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Generate new access token
	token, err := utils.GenerateToken(user.Email, time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("could not generate token: %v", err)
	}

	return &authpb.AuthResponse{
		Message: "Access token refreshed",
		Token:   token,
		RefreshToken: rtoken.Token,
	}, nil
}
