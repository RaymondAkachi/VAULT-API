package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {


	list, err := net.Listen("tcp", fmt.Sprintf(":d%", 9000))
	if err != nil {
		log.Fatalf("Server refused to listen becuase of this reason: %v", err)
	}


	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %s", err)
	}
}