package main

import (
	"fmt"
	"log"
	"net"

	"product_svc/pkg/config"
	"product_svc/pkg/di"
	"product_svc/pkg/product/pb"
	"product_svc/pkg/verification"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error while loading config", err)
	}

	verification.InitTwilio(config)

	service, _ := di.InitializeAPI(config)

	lis, err := net.Listen("tcp", config.ProductSvcUrl)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	fmt.Println("Auth Svc on", config.ProductSvcUrl)

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
