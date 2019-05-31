package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/microhq/blob-srv/handler"

	pb "github.com/microhq/blob-srv/proto/blob"
)

var (
	buffSize int
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.blob"),
		micro.Version("latest"),
		micro.Flags(
			cli.IntFlag{
				Name:  "buf_size",
				Value: 1024 * 1024,
				Usage: "Download buffer size",
			},
		),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			buffSize = c.Int("buf_size")
		}),
	)

	// craete Blob service handler
	blob, err := handler.NewBlob(".", buffSize)
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
