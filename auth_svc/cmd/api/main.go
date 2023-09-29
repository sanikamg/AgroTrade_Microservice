package main

import (
	"auth_svc/pkg/auth/pb"
	"auth_svc/pkg/config"
	"auth_svc/pkg/di"
	"auth_svc/pkg/verification"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error while loading config", err)
	}

	verification.InitTwilio(config)

	service := di.InitializeAPI(config)

	lis, err := net.Listen("tcp", config.AuthSvcUrl)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	fmt.Println("Auth Svc on", config.AuthSvcUrl)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
