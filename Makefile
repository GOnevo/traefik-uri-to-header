.PHONY: all fmt test cover yaegi_test dep

all: fmt test yaegi_test

fmt:
	gofmt -s -w ./
test: ## Run data race detector
	go test -v ./...
cover: ## Run unittests
	go test -cover -short -coverprofile=cover.out ./...
	go tool cover -html=cover.out
yaegi_test:
	yaegi test -v .
dep: ## Get the dependencies
	go mod download

.DEFAULT_GOAL := all
