package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sergiommarcial/gen" // Import your generated proto package

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	return &pb.DataResponse{Reply: "Hello, " + req.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterMyServiceServer(server, pb.UnimplementedMyServiceServer{})
	log.Println("Server started on :50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
