PROJECT_NAME=3rd-party-gateway
BUILD_VERSION=1.0.0
GITHUB_USERNAME=username

DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

compose_dev: docker
	docker tag $(PROJECT_NAME):$(BUILD_VERSION) $(GITHUB_USERNAME)/$(PROJECT_NAME):$(BUILD_VERSION); \
	docker push $(GITHUB_USERNAME)/$(PROJECT_NAME):$(BUILD_VERSION);
	cd deploy/; \
	docker-compose up -d;

build:
	$(GO_BUILD_ENV) go build -v -o $(PROJECT_NAME)-$(BUILD_VERSION).bin main.go

docker_prebuild: build
	mv $(PROJECT_NAME)-$(BUILD_VERSION).bin ./$(PROJECT_NAME).bin; \

docker_build:
	cd ./; \
	docker build --rm -t $(DOCKER_IMAGE) .;

docker_postbuild:
	rm -rf ./$(PROJECT_NAME).bin 2> /dev/null;\

docker: docker_prebuild docker_build docker_postbuild