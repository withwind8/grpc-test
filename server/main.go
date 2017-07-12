package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpctest "github.com/withwind8/grpc-test/proto"
)

type server struct{}

func (s *server) Sum(ctx context.Context, in *grpctest.MathReq) (*grpctest.MathResp, error) {
	log.Printf("RPC call: %d + %d ", in.N1, in.N2)
	return &grpctest.MathResp{Sum: in.N1 + in.N2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	grpctest.RegisterMathServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
