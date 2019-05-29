# Blob Service

This is an example of a simple file blob microservice.

Generated with

```
micro new blob-srv --namespace=go.micro --type=srv --gopath=false
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)
- [Example](#example)

## Configuration

- FQDN: go.micro.srv.blob
- Type: srv
- Alias: blob

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./blob-srv
```

Build a docker image
```
make docker
```

## Example

The project provides examples of simple upload and download clients in `examples` directory. See the example below how to run them:

Run the service:
```shell
MICRO_REGISTRY=consul go run main.go
2019/05/29 21:50:23 Transport [http] Listening on [::]:64684
2019/05/29 21:50:23 Broker [http] Connected to [::]:64685
2019/05/29 21:50:23 Registry [consul] Registering node: go.micro.srv.blob-b78bae8e-e553-414f-a417-1a2f3014634c
2019/05/29 21:50:26 Received request to create bucket: foobar
2019/05/29 21:50:26 Received PUT request
2019/05/29 21:50:26 Received 1048576 bytes of blob: foobar/foobar.out
2019/05/29 21:50:26 Received 1048576 bytes of blob: foobar/foobar.out
...
...
2019/05/29 21:50:26 Received 1048576 bytes of blob: foobar/foobar.out
2019/05/29 21:50:26 Received 843332 bytes of blob: foobar/foobar.out
2019/05/29 21:50:26 Finished receiving data
2019/05/29 21:55:39 Received GET request to get file foobar.out from bucket foobar
2019/05/29 21:55:39 Sending 1048576 bytes of foobar/foobar.out
2019/05/29 21:55:39 Sending 1048576 bytes of foobar/foobar.out
...
...
2019/05/29 21:55:39 Sending 843332 bytes of foobar/foobar.out
```

Run the uploader:
```shell
MICRO_REGISTRY=consul go run upload.go --bucket_id="foobar" --blob_path="foobar.out"
2019/05/29 21:50:26 streaming 1048576 bytes for blob: foobar.out
2019/05/29 21:50:26 streaming 1048576 bytes for blob: foobar.out
...
...
2019/05/29 21:50:26 streaming 843332 bytes for blob: foobar.out
```

Run the downloader:
```shell
MICRO_REGISTRY=consul go run download.go --bucket_id="foobar" --key_id="foobar.out" --file_path="foobar.out"
2019/05/29 21:55:39 received 1048576 bytes for blob: foobar.out
2019/05/29 21:55:39 received 1048576 bytes for blob: foobar.out
...
...
2019/05/29 21:55:39 received 843332 bytes for blob: foobar.out
```
