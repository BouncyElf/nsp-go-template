package main

import (
	"flag"
	"log"
	"net"
	"nsp-go-template/conf"
	"nsp-go-template/pb"
	"nsp-go-template/service"

	"google.golang.org/grpc"
)

var (
	configPath = flag.String("c", "", "config file path")
)

func main() {
	flag.Parse()

	err := conf.ParseConfigFromFile(*configPath)
	if err != nil {
		log.Fatalf("load config file error: %v", err)
	}

	lis, err := net.Listen(conf.ServiceConf.Network, conf.ServiceConf.Addr)
	if err != nil {
		log.Fatalf("error listen to %s network:%s error:%v", conf.ServiceConf.Addr, conf.ServiceConf.Network, err)
	}

	server := grpc.NewServer()
	pb.RegisterValleyServer(server, new(service.ValleyService))

	if err := server.Serve(lis); err != nil {
		log.Fatalf("error Serve: %v", err)
	}
}
