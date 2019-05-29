GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=.:${GOPATH}/src --micro_out=. --go_out=. proto/blob/blob.proto

.PHONY: build
build: proto

	go build -o blob-srv main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t microhq/blob-srv:latest
