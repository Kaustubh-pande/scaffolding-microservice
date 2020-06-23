GOPATH:=$(shell go env GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
deps:
	$(GOGET) golang.org/x/tools/cmd/stringer
	$(GOGET) "google.golang.org/grpc"
	$(GOGET) "github.com/grpc-ecosystem/grpc-gateway/runtime"
	$(GOGET) "github.com/golang/protobuf/jsonpb"
	$(GOGET) "google.golang.org/grpc/reflection"
build:
	$(GOBUILD) -o bin/main main.go
