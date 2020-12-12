osrm-dowload: ## Download data for local osrm instance
	cd ./data && ../scripts/osrm_download.sh

osrm-launch: ## Run local osrm instance
	cd ./data && ../scripts/osrm_launch.sh

run: ## Rebuild and run web service
	cd ./app && go build && ./ingrid-coding-assignment serve

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
