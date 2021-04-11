package main

import (
	"product/handler"
	pb "product/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("product"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterProductHandler(srv.Server(), new(handler.Product))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
