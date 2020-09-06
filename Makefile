OS = Linux
VERSION = 0.0.1
MODULE = amf

# git commit id
COMMITID ?= latest
# Image URL to use all building/pushing image targets
IMG ?= hub.docker.com/uhhc/${MODULE}:${COMMITID}
IMG_BASE ?= hub.docker.com/uhhc/amf-base:${COMMITID}

ROOT_PACKAGE=github.com/uhhc/${MODULE}
CURDIR = $(shell pwd)
SOURCEDIR = $(CURDIR)
COVER = $($3)

ECHO = echo
RM = rm -rf
MKDIR = mkdir

.PHONY: test

default: test lint vet

test:
	go test -cover=true $(PACKAGES)

race:
	go test -cover=true -race $(PACKAGES)

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

# https://godoc.org/golang.org/x/tools/cmd/goimports
# imports:
# 	goimports ./...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./...

all: test

PACKAGES = $(shell go list ./... | grep -v './vendor/\|./tests\|./mock')
BUILD_PATH = $(shell if [ "$(CI_DEST_DIR)" != "" ]; then echo "$(CI_DEST_DIR)" ; else echo "$(PWD)"; fi)

cover: collect-cover-data test-cover-html open-cover-html

collect-cover-data:
	echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES),\
		go test -v -coverprofile=coverage.out -covermode=count $(pkg);\
		if [ -f coverage.out ]; then\
			tail -n +2 coverage.out >> coverage-all.out;\
		fi;)

test-cover-html:
	go tool cover -html=coverage-all.out -o coverage.html

test-cover-func:
	go tool cover -func=coverage-all.out

open-cover-html:
	open coverage.html

build-local:
	@$(ECHO) "Will build on "$(BUILD_PATH)
	go mod vendor
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -a -ldflags "-w -s" -v -o $(BUILD_PATH)/bin/${MODULE} $(ROOT_PACKAGE)

build:
	@$(ECHO) "Will build on "$(BUILD_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags "-w -s" -v -o $(BUILD_PATH)/bin/${MODULE} $(ROOT_PACKAGE)

run: fmt lint
	go run main.go serve

clean:
	rm -f *.out *.html

compile: test build

docker-build:
	docker build -t ${IMG} .

service-reload:
	docker-compose -p amf -f devbox/${MODULE}.yaml stop
	docker-compose -p amf -f devbox/${MODULE}.yaml up -d --remove-orphans

mysql-reload:
	docker-compose -p mysql -f devbox/mysql.yaml stop
	docker-compose -p mysql -f devbox/mysql.yaml up -d --remove-orphans

mongo-reload:
	docker-compose -p mongo -f devbox/mongodb.yaml stop
	docker-compose -p mongo -f devbox/mongodb.yaml up -d --remove-orphans

swagger:
	swag init --parseDependency -o ./swagger_docs
	mv ./swagger_docs/swagger.json assets/swagger-ui
	rm -rf ./swagger_docs

grpc:
	# Format: protoc -I <proto_import_path1> -I <proto_import_path2> --go_out=<plugins={plugin1+plugin2+...}>:<pb_output_path> <proto_file_path>
# 	protoc -I pkg/grpc/proto --gogofaster_out=plugins=grpc:pkg/grpc/pb pkg/grpc/proto/*.proto
	protoc -I pkg/grpc/proto --go_out=plugins=grpc:pkg/grpc/pb pkg/grpc/proto/*.proto

help:
	@$(ECHO) "Targets:"
	@$(ECHO) "all				- test"
	@$(ECHO) "setup				- install necessary libraries"
	@$(ECHO) "test				- run all unit tests"
	@$(ECHO) "cover [package]	- generates and opens unit test coverage report for a package"
	@$(ECHO) "race				- run all unit tests in race condition"
	@$(ECHO) "add				- run govendor add +external command"
	@$(ECHO) "build-local		- build and exports locally"
	@$(ECHO) "build				- build and exports using CI_DEST_DIR"
	@$(ECHO) "run				- run the program"
	@$(ECHO) "clean				- remove test reports and compiled package from this folder"
	@$(ECHO) "compile			- test and build - one command for CI"
	@$(ECHO) "docker-build		- builds an image with this folder's Dockerfile"
	@$(ECHO) "service-reload	- run service docker compose"
	@$(ECHO) "mysql-reload		- run a mysql docker compose"
	@$(ECHO) "mongo-reload		- run a mongodb docker compose"
	@$(ECHO) "swagger		    - init or rebuild swagger.json"
	@$(ECHO) "grpc		        - init or rebuild protocol buffers files"