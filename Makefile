

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Buid application
	go get -u
	go mod vendor
	go build -o bin/api/main.exe cmd/api/main.go

run: ## Run application
	./bin/api/main.exe

test:  ## Test application
	go test ./...
	