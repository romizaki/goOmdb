// main.go

package main

import (
	"log"
	"net"

	pb "omdb/internal/types"

	omdb_service "omdb/internal/omdb.service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOMDBServiceServer(grpcServer, &omdb_service.OMDBServer{})

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
