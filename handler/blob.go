package handler

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/micro/go-log"
	pb "github.com/microhq/blob-srv/proto/blob"
)

// Blob is a file blob handler
type Blob struct {
	// baseDir is the base directory under which all blobs are stored
	baseDir string
}

// NewBlob creates new blob handler and returns it
func NewBlob(dir string) (*Blob, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("path does not exist: %s", dir)
	}

	return &Blob{baseDir: dir}, nil
}

// CreateBucket creates new bucket to store the files blobs in.
// Bucket can be specified as a "/" separated list of names which create a bucket hierarchy.
func (b *Blob) CreateBucket(ctx context.Context, req *pb.CreateBucketReq, resp *pb.CreateBucketResp) error {
	log.Logf("Received request to create bucket: %s", req.Id)

	return os.MkdirAll(filepath.Join(b.baseDir, req.Id), 0755)
}

// DeleteBucket deletes bucket.
func (b *Blob) DeleteBucket(ctx context.Context, req *pb.DeleteBucketReq, resp *pb.DeleteBucketResp) error {
	log.Logf("Received request to delete bucket: %s", req.Id)

	return os.Remove(filepath.Join(b.baseDir, req.Id))
}

// Put allows to upload a file into specified bucket.
// It fails with error if either the client upload fails or the file fails to be stored.
func (b *Blob) Put(ctx context.Context, stream pb.Blob_PutStream) error {
	log.Logf("Received PUT request")

	var id string
	var bucketId string
	defer stream.Close()

	f, err := ioutil.TempFile(b.baseDir, "example")
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		blob, err := stream.Recv()
		if err == io.EOF {
			log.Logf("Finished receiving data")
			break
		}

		if err != nil {
			return fmt.Errorf("Error storing blob: %s", err)
		}

		bucketId = blob.BucketId
		id = blob.Id

		log.Logf("Received %d bytes of blob: %s", len(blob.Data), filepath.Join(blob.BucketId, blob.Id))

		if _, err := f.Write(blob.Data); err != nil {
			return err
		}
	}

	oldPath := filepath.Join(b.baseDir, f.Name())
	newPath := filepath.Join(b.baseDir, filepath.Join(bucketId, id))

	// move tempfile into bucket directory
	if err := os.Rename(oldPath, newPath); err != nil {
		log.Logf("Error renaming files: %s", err)
		return err
	}

	return nil
}

// Get allows to download a file from given bucket.
// Get streans the requested file to client in chunks of 1MB.
func (b *Blob) Get(ctx context.Context, req *pb.GetReq, stream pb.Blob_GetStream) error {
	log.Logf("Received GET request to get file %s from bucket %s", req.Id, req.BucketId)

	filePath := filepath.Join(b.baseDir, req.BucketId, req.Id)
	file, err := os.Open(filePath)
	if err != nil {
		log.Logf("failed to open file %s: %s", filePath, err)
		return err
	}
	defer file.Close()

	defer stream.Close()
	buf := make([]byte, 1024*1024)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("Error streaming file: %s", err)
		}

		log.Logf("Sending %d bytes of %s", n, filepath.Join(req.BucketId, req.Id))

		if err := stream.Send(&pb.GetResp{Data: buf[:n]}); err != nil {
			return fmt.Errorf("Error streaming file %s: %s", file.Name(), err)
		}
	}

	return nil
}

// Delete deletes the file given by its id from the given bucket.
// It returns error of the file failed to be deleted from blob store.
func (b *Blob) Delete(ctx context.Context, req *pb.DeleteReq, resp *pb.DeleteResp) error {
	log.Logf("Received DELETE request for file %s", filepath.Join(req.BucketId, req.Id))

	return os.Remove(filepath.Join(b.baseDir, req.BucketId, req.Id))
}

// List lists all the keys inside the given bucket.
func (b *Blob) List(ctx context.Context, req *pb.ListReq, resp *pb.ListResp) error {
	log.Logf("Received LIST request for bucket: %s", req.BucketId)

	dir, err := os.Open(filepath.Join(b.baseDir, req.BucketId))
	if err != nil {
		log.Logf("failed to open bucket %s: %s", req.BucketId, err)
		return err
	}
	defer dir.Close()

	keys, err := dir.Readdirnames(0)
	if err != nil {
		log.Logf("failed to list keys in bucket %s: %s", req.BucketId, err)
		return err
	}

	resp.Id = keys

	return nil
}
