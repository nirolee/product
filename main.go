package main

import (
	"github.com/nirolee/product.git/handler"
	pb "github.com/nirolee/product.git/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("github.com/nirolee/product.git"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGithub.Com/Nirolee/Product.GitHandler(srv.Server(), new(handler.Github.Com/Nirolee/Product.Git))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
