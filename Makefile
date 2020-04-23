tag := latest
server_in := cmd/api/main.go
server_out := main

.PHONY: default
default: docker

.PHONY: ci
ci: docker push ## Build in CI

.PHONY: docker
docker: ## Build Docker image

	docker build --progress=plain \
		--tag $(service):latest \
		--tag $(service):$(tag) \

.PHONY: push
push: ## Push images to registry
	docker push $(service):latest
	docker push $(service):$(tag)

.PHONY: server
server: dep ## Compile server for local OS
	go build -i -v -o $(server_out) $(server_in)

.PHONY: up
up: ## Starts go-example-api
	-docker network create example-go-network
	docker-compose up

.PHONY: down
down: ## Shutsdown the server
	docker-compose down

.PHONY: logs
logs: ## Tails the logs on the Docker container
	docker-compose logs -f

.PHONY: test
test: ## Runs any tests in the current directory tree
	go test -v -cover ./...

.PHONY: clean
clean: ## Remove previous builds
	rm -rfv $(server_out)
	rm -f dev_docker
	docker system prune -f
	docker volume prune -f

.PHONY: dep
dep: ## Get the dependencies
	go get -v -d ./...

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
