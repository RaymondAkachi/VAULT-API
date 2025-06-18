package main

import (
	"log"
	"net"

	authpb "github.com/RaymondAkachi/VAULT-API/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	InitDB()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authServer := &AuthServer{}

	authpb.RegisterAuthServiceServer(grpcServer, authServer)

	// This enables reflection for Evans and other tools
	reflection.Register(grpcServer)

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}