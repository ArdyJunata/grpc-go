package main

import (
	"log"
	"net"

	"github.com/ArdyJunata/grpc-go/apps/auth"
	"github.com/ArdyJunata/grpc-go/external/database"
	"github.com/ArdyJunata/grpc-go/internal/config"
	"google.golang.org/grpc"
)

func main() {
	err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", config.Cfg.App.AuthPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RouterInitGRPC(s, db)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
