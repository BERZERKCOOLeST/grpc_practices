package main

import (
	"grpc_soldshop/src/proto/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const port = ":50051"

func main() {

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("../auth/server/server.pem", "../auth/server/server.key")
	if err != nil {
		log.Fatalf("failed to create creds: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
