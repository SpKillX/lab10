package main

import (
	"context"
	"net"
	pb "task_1/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBookingServiceServer
}

func (s *server) CheckTable(ctx context.Context, in *pb.TableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{Available: true}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterBookingServiceServer(s, &server{})
	s.Serve(lis)
}
