all: lint test build start

### Development

start:
	./bin/InfobaeAPI

dev: path
	gowatch

fmt:
	@if [ -n "$$(go fmt ./src)" ]; then \
        echo "Code is not properly formatted"; \
        exit 1; \
    fi

lint:
	golangci-lint run

clear:
	rm -rf ./bin

## https://github.com/golang-standards/project-layout/issues/113#issuecomment-1336514449
build: clear fmt
	GOARCH=amd64 go build -o ./bin/InfobaeAPI ./src/main.go

build-arm: clear fmt
	GOARCH=arm64 go build -o ./bin/InfobaeAPI ./src/main.go

test:
	go test -v ./tests

path:
	export PATH=$PATH:$HOME/go/bin

### Kubernetes && Docker

## https://www.digitalocean.com/community/tutorials/how-to-use-minikube-for-local-kubernetes-development-and-testing
k8s-up:
	minikube start

k8s-down:
	minikube stop

k8s-apply:
	kubectl apply -f ./kubernetes

pods:
	kubectl get pods -A

compose-up:
	docker compose -f ./docker/docker-compose.yml up -d --build

compose-down:
	docker compose -f ./docker/docker-compose.yml down