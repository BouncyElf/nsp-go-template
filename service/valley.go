package service

import (
	"context"
	"nsp-go-template/pb"
)

type ValleyService struct {
	pb.UnimplementedValleyServer
}

func (*ValleyService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Echo: req.GetSay(),
	}, nil
}
