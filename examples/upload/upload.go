package main

import (
	"context"
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
		path     string
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
				Name:  "blob_path",
				Value: "",
				Usage: "Path to file blob",
			},
		),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			bucketId = c.String("bucket_id")
			path = c.String("blob_path")
		}),
	)

	// Initialize service client
	client := pb.NewBlobService("go.micro.srv.blob", service.Client())

	// Create bucket with given bucket ID
	if _, err := client.CreateBucket(context.Background(), &pb.CreateBucketReq{Id: bucketId}); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Logf("failed to open file %s", path)
		return
	}
	defer file.Close()

	// request stream object to upload the blob via
	stream, err := client.Put(context.Background())
	if err != nil {
		log.Logf("failed to create upload file %s", path)
		return
	}
	defer stream.Close()

	// flag to control upload
	upload := true

	// sent data in 1M chunks
	buf := make([]byte, 1024*1024)

	// stream the data across to blob service
	for upload {
		n, err := file.Read(buf)
		if err == io.EOF {
			upload = false
			break
		}

		if err != nil {
			log.Logf("error reading file %s: %s", path, err)
			break
		}

		log.Logf("streaming %d bytes for blob: %s", n, path)

		if err := stream.Send(&pb.PutReq{Id: path, BucketId: bucketId, Data: buf[:n]}); err != nil {
			log.Logf("error streaming blob %s: %s", path, err)
			break
		}
	}
}
