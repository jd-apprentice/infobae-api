all: lint test imports swag build start

### Development

start:
	./bin/InfobaeAPI

dev: docs path
	air -c .air.toml

fmt:
	@if [ -n "$$(go fmt ./src)" ]; then \
        echo "Code is not properly formatted"; \
        exit 1; \
    fi

fmt-fix:
	go fmt ./src

imports: path
	goimports -w ./src

lint: path
	golangci-lint run --exclude-dirs=docs

lint-fix: path
	golangci-lint run --fix --exclude-dirs=docs

docs: path
	swag init -d src

clear:
	rm -rf ./bin

cache:
	go clean -modcache

## https://github.com/golang-standards/project-layout/issues/113#issuecomment-1336514449
build: clear fmt
	GOARCH=amd64 go build -o ./bin/InfobaeAPI ./src/main.go

build-arm: clear fmt
	GOARCH=arm64 go build -o ./bin/InfobaeAPI ./src/main.go

test:
	go test -v ./tests

path:
	export PATH=$$PATH:$$HOME/go/bin

### Kubernetes && Docker

## https://www.digitalocean.com/community/tutorials/how-to-use-minikube-for-local-kubernetes-development-and-testing
k8s-up:
	minikube start

k8s-down:
	minikube stop

k8s-apply:
	kubectl apply -f ./infra/kubernetes

pods:
	kubectl get pods -A

compose-up:
	docker compose -f ./infra/docker/docker-compose.yml up -d --build 

compose-down:
	docker compose -f ./infra/docker/docker-compose.yml down