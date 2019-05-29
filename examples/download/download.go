package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	pb "github.com/microhq/blob-srv/proto/blob"
)

func main() {
	var (
		bucketId string
		keyId    string
		filePath string
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.blob"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:  "bucket_id",
				Value: "",
				Usage: "Bucket name",
			},
			cli.StringFlag{
				Name:  "key_id",
				Value: "",
				Usage: "Blob key id",
			},
			cli.StringFlag{
				Name:  "file_path",
				Value: "",
				Usage: "Path to store the blbo in",
			},
		),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			bucketId = c.String("bucket_id")
			keyId = c.String("key_id")
			filePath = c.String("file_path")
		}),
	)

	// Initialize service client
	client := pb.NewBlobService("go.micro.srv.blob", service.Client())

	// List all available keys in given bucket
	keys, err := client.List(context.Background(), &pb.ListReq{BucketId: bucketId})
	if err != nil {
		log.Logf("failed to list bucket %s: %s", bucketId, err)
	}

	for _, key := range keys.Id {
		log.Logf("key: %s", key)
	}

	// request stream to use for blob download
	stream, err := client.Get(context.Background(), &pb.GetReq{Id: keyId, BucketId: bucketId})
	if err != nil {
		log.Logf("failed to get download file: %s", err)
		return
	}
	defer stream.Close()

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0640)
	if err != nil {
		log.Logf("failed to open file %s", filePath)
		return
	}
	defer file.Close()

	// flag to control download progress
	download := true

	for download {
		blob, err := stream.Recv()
		if err == io.EOF {
			log.Logf("finished receiving data")
			break
		}

		if err != nil {
			fmt.Errorf("error downloading %s from bucket %s: %s", keyId, bucketId, err)
			break
		}

		log.Logf("received %d bytes for blob: %s", len(blob.Data), keyId)

		if _, err := file.Write(blob.Data); err != nil {
			fmt.Errorf("error writing data to file %s: %s", file.Name(), err)
			break
		}
	}
}
