package main

import (
	"log"
	"net"

	"github.com/ksntrvsk/grpc_nmap_wrapper/api/pb"
	"github.com/ksntrvsk/grpc_nmap_wrapper/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	srv := &server.GRPCServer{}
	pb.RegisterNetVulnServiceServer(grpcServer, srv)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
