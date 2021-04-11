package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	github.com/nirolee/product.git "github.com/nirolee/product.git/proto"
)

type Github.Com/Nirolee/Product.Git struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Github.Com/Nirolee/Product.Git) Call(ctx context.Context, req *github.com/nirolee/product.git.Request, rsp *github.com/nirolee/product.git.Response) error {
	log.Info("Received Github.Com/Nirolee/Product.Git.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Github.Com/Nirolee/Product.Git) Stream(ctx context.Context, req *github.com/nirolee/product.git.StreamingRequest, stream github.com/nirolee/product.git.Github.Com/Nirolee/Product.Git_StreamStream) error {
	log.Infof("Received Github.Com/Nirolee/Product.Git.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&github.com/nirolee/product.git.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Github.Com/Nirolee/Product.Git) PingPong(ctx context.Context, stream github.com/nirolee/product.git.Github.Com/Nirolee/Product.Git_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&github.com/nirolee/product.git.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
