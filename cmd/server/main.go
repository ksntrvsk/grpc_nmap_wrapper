package main

import (
	"fmt"
	"net"

	"github.com/ksntrvsk/grpc_nmap_wrapper/api/pb"
	"github.com/ksntrvsk/grpc_nmap_wrapper/pkg/config"
	"github.com/ksntrvsk/grpc_nmap_wrapper/pkg/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	// Constants
	const (
		network = "tcp"
	)

	// Create grpcServer
	grpcServer := grpc.NewServer()
	srv := &server.GRPCServer{}
	pb.RegisterNetVulnServiceServer(grpcServer, srv)

	// Create config file
	cfg, err := config.NewCongif()
	if err != nil {
		log.Fatalf("unable to get a config: %v", err)
	}

	// Set logging level
	level, err := log.ParseLevel(cfg.Logger.Level)
	if err != nil {
		log.Errorf("unable to parse level: %v", err)
	}
	log.SetLevel(level)

	// Create listener
	listener, err := net.Listen(
		network,
		fmt.Sprintf(":%s", cfg.Server.Port),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server start")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
