package middleware

import (
	"context"
	"errors"
	"strings"

	utils "github.com/RaymondAkachi/VAULT-API/auth-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}

		authHeaders := md["authorization"]
		if len(authHeaders) == 0 {
			return nil, errors.New("authorization header not provided")
		}

		// Except format: Bearer <token>
		tokenStr := strings.TrimPrefix(authHeaders[0], "Bearer ")
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			return nil, errors.New("invalid or expired token")
		}

		//Add claims to context for downstrram handlers
		ctx = context.WithValue(ctx, "claims", claims)

		//Continue executon
		return handler(ctx, req)



	}
}