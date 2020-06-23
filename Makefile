GOPATH:=$(shell go env GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build

build:
	$(GOBUILD) -o bin/main main.go
