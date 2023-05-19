APP=goHttpServerGin

.PHONY: help all build windows linux

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    help   : Show help"
	@echo "    all    : Build multiple binary of this project"
	@echo "    build  : Build the binary of this project for current platform"
	@echo "    windows: Build the windows binary of this project"
	@echo "    linux  : Build the linux binary of this project"
all:build windows linux
build:
	@go build -o ${APP}
windows:
	@go env -w GOOS=windows
	@go env -w GOARCH=amd64
	@go build -o ${APP}-windows-amd64
linux-arm:
	@go env -w GOOS=linux
	@go env -w GOARCH=arm
	@go build -o ${APP}-linux-arm
linux-64:
	@go env -w GOOS=linux
	@go env -w GOARCH=amd64
	@go build -o ${APP}-linux-64