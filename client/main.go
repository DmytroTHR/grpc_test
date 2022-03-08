package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"svcpod-client/proto"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	servicePort = "8989"
	httpPort    = "8888"
)

var serviceName = envOrDefault("MyServiceName", "localhost")

func envOrDefault(envName, defaultValue string) string {
	res := os.Getenv(envName)
	if res == "" {
		return defaultValue
	}
	return res
}

func getMyService() *grpc.ClientConn {
	grpcServer := net.JoinHostPort(serviceName, servicePort)
	myConnection, err := grpc.Dial(grpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("%s: unable to set grpc connection - %v", grpcServer, err)
	}
	return myConnection
}
func main() {

	myConnection := getMyService()
	defer myConnection.Close()

	httpServer := &http.Server{
		Handler: NewRequestHandler(proto.NewSvcPodServClient(myConnection)),
		Addr:    ":" + httpPort,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		log.Println("Server shutting down..")
		close(idleConnsClosed)
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
