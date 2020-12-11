osrm-dowload: ## Run local osrm instance
	cd ./data && ../scripts/osrm_download.sh

osrm-launch: ## Run local osrm instance
	cd ./data && ../scripts/osrm_launch.sh

# Make makefile autodocumented
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
