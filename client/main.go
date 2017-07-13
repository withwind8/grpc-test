package main

import (
	"context"
	"log"
	"os"
	"strconv"

	grpctest "github.com/withwind8/grpc-test/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := grpctest.NewMathClient(conn)

	n1 := 1
	n2 := 2
	if len(os.Args) > 2 {
		n1, _ = strconv.Atoi(os.Args[1])
		n2, _ = strconv.Atoi(os.Args[2])
	}

	r, err := c.Jian(context.Background(), &grpctest.MathReq{int32(n1), int32(n2)})
	if err != nil {
		log.Fatalf("rpc fail: %v", err)
	}
	log.Printf("RPC: %d", r.Sum)
}
