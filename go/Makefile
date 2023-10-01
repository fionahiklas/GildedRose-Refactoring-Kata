.PHONY: help tidy generate install_tools lint format test test_clean test_report test_summary

# from https://suva.sh/posts/well-documented-makefiles/
# sorting from https://stackoverflow.com/questions/14562423/is-there-a-way-to-ignore-header-lines-in-a-unix-sort
# Add double hash comments after every target to provide help text
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) | awk 'NR<6{print $0; next}{print $0 | "sort"}'

version: ## Simply output the version number, used by some build scripts
	@echo "${CODE_VERSION}"

clean: ## Clean temporary/build files
	rm -f coverage.out
	find build -not -path '*/.*' -a -type f -print | xargs rm -f

tidy: ## Tidy imports
	go mod tidy

download: ## Ensure modules are downloaded
	go mod download
	go mod verify

install_tools: download ## Install tooling needed by other targets
	CGO_ENABLED=0 go install golang.org/x/tools/cmd/goimports
	CGO_ENABLED=0 go install github.com/golang/mock/mockgen
	# This tool actually uses the cgo support and shared libraries :/
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: ## Lint codebase
	golangci-lint -E goimports run ./...

format: ## Format codebase and check imports
	goimports -w ./*.go gildedrose

delete_mock: ## Deletes all files that match `mock_*.go`
	@echo "Deleting all mocks"
	find . -name "mock_*.go" -print -delete

generate: delete_mock ## Generate mocks
	@echo "Generating mocks ..."
	go generate ./internal/... ./pkg/...

clean_test_cache: ## Allows all tests to be forced to run without using cached results
	go clean -testcache

test_clean: clean_test_cache test ## Run tests but clear the cache first

test_report: test ## Run unit tests and generate an HTML coverage report
	go tool cover -html=coverage.out

test_summary: test ## Output coverage summary
	go tool cover -func=coverage.out

test: ## Run unit tests and check coverage
	@echo "Running unit tests"
	go test ${GO_LD_FLAGS} -coverprofile=coverage.out ./...

run: ## Run the test fixture
	go run ./ ${DAYS}