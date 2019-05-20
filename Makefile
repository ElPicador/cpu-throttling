all: build docker push-docker

build:
	@GOBIN="cpu-throttling" GOOS=linux GOARCH=amd64 go build --ldflags '-extldflags "-static"'

docker:
	@docker build -t elpicador/cpu-throttling .

push-docker:
	@docker push elpicador/cpu-throttling:latest

.PHONY: build docker push-docker all
