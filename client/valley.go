package main

import (
	"context"
	"fmt"
	"log"
	"nsp-go-template/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:4321",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewValleyClient(conn)
	resp, err := c.Echo(context.Background(), &pb.EchoRequest{
		Say: "hi, nice to see u",
	})
	if err != nil {
		log.Fatalf("call Echo error: %v", err)
	}
	fmt.Println(resp)
}
