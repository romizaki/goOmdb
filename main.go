// main.go

package main

import (
	"log"
	"net"

	omdb_service "omdb/internal/omdb.service"
	pb "omdb/internal/types"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	omdbServer := &omdb_service.OMDBServer{}
	pb.RegisterOMDBServiceServer(grpcServer, omdbServer)

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
