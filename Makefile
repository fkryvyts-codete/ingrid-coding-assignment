osrm-dowload: ## Download data for local osrm instance
	cd ./data && ../scripts/osrm_download.sh

osrm-launch: ## Run local osrm instance
	cd ./data && ../scripts/osrm_launch.sh

test: ## Run tests
	cd ./app && go test -v ./.../service

build: ## Rebuild the web service without using docker
	cd ./app && go mod download && go build

run: ## Run the web service without using docker
	cd ./app && ./ingrid-coding-assignment serve

build-docker: ## Rebuild the web service using docker
	cd ./app && docker build -t ingrid-coding-assignment -f Dockerfile .
run-docker: ## Run the web service using docker
	docker run -it -p 8080:8080 ingrid-coding-assignment
push-docker: ## Push built docker image to Docker Hub
	docker tag ingrid-coding-assignment fkryvyts/ingrid-coding-assignment
	docker login
	docker push fkryvyts/ingrid-coding-assignment

lint: ## Run linter for the code
	cd ./app && golangci-lint run

lint-fix: ## Run linters and perform automatic fixes
	cd ./app && golangci-lint run --fix

pre-commit: ## Install pre-commit hooks
	pre-commit install

# Make makefile autodocumented
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
