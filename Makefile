PKG=$(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION=v$(shell cat .version)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X ${PKG}/version.Version=${VERSION}"
OPENAPI=tools openapi

Workdir ?= ./cmd/srv-go-blog

up: dockerize
	cd $(Workdir) && go run .

dockerize:
	cd $(Workdir) && go run . dockerize

migrate:
	cd $(Workdir) && go run . migrate

release:
	git push
	git push origin ${VERSION}
