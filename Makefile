# Makefile for nudex_voter

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=nudex-voter

# Build binary
all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

deps:
	$(GOGET) -u ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd && ./$(BINARY_NAME)

docker-build-all:
	docker buildx build --platform linux/amd64,linux/arm64 -t nuvosphere/nudex-voter:latest --push .

docker-build:
	docker buildx build --platform linux/amd64 -t nuvosphere/nudex-voter:latest --load .

docker-build-x:
	docker buildx build --platform linux/arm64 -t nuvosphere/nudex-voter:latest --load .

.PHONY: fmt
fmt: ## go fmt
	go fmt ./...
	#gofumpt -l -w .

.PHONY: tool
tool: ## Install dep tool
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/4meepo/tagalign/cmd/tagalign@latest

.PHONY: abi
abi: ## generate abi go file
	abigen --abi internal/layer2/contracts/AccountManager.json --pkg abis --type AccountManagerContract --out internal/layer2/abis/account_manager.go
	abigen --abi internal/layer2/contracts/NuDexOperations.json --pkg abis --type NuDexOperationsContract --out internal/layer2/abis/nudex_operations.go
	abigen --abi internal/layer2/contracts/ParticipantManager.json --pkg abis --type ParticipantManagerContract --out internal/layer2/abis/participant_manager.go
	abigen --abi internal/layer2/contracts/VotingManager.json --pkg abis --type VotingManagerContract --out internal/layer2/abis/voting_manager.go

.PHONY: help
help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	
.PHONY: all build clean test deps run docker-build