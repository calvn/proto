GOOS?=linux
export GOOS

build:
	@echo " ==> Building $(GOOS) executable..."
	@go build -o build/proto .
	@echo " ==> Building Docker image..."
	@docker build -t calvn/proto .
.PHONY: build

push:
	@echo " ==> Pushing image to Docker registry..."
	@docker push calvn/proto
.PHONY: push
