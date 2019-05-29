package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/microhq/blob-srv/handler"

	pb "github.com/microhq/blob-srv/proto/blob"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.blob"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// craete Blob service handler
	blob, err := handler.NewBlob(".")
	if err != nil {
		log.Fatal(err)
	}

	// Register Handler
	pb.RegisterBlobHandler(service.Server(), blob)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
