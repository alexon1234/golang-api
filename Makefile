

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Buid functions
	go get -u
	go mod vendor
	go build ./...
	done


test:  ## Test functions
	go test ./...

deploy: ## Deploy functions
	sls deploy
	