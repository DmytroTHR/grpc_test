package main

import (
	"log"
	"net"
	"svcpod-server/proto"
	"svcpod-server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	serviceName = "svcpod-server"
	servicePort = "8989"
)

func main() {
	service := service.NewMyService()
	listener, err := net.Listen("tcp", net.JoinHostPort("", servicePort))
	if err != nil {
		log.Panicf("%s: failed to listen on port - %v", serviceName, err)
	}

	server := grpc.NewServer()
	defer server.GracefulStop()
	proto.RegisterSvcPodServServer(server, service)
	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Panicf("%s: failed to start grpc - %v", serviceName, err)
	}
}
