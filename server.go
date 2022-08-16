package main

import (
	"fmt"
	"gserver/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world")
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	keepaliveOpts := grpc.MaxRecvMsgSize(1024 * 1024 * 4)
	grpcServer := grpc.NewServer(keepaliveOpts)

	s := api.Server{}

	api.RegisterStudentServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server %v", err)
	}

}

