# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

deps-audit: ### check dependencies vulnerabilities
	govulncheck ./...
.PHONY: deps-audit

format: ### Run code formatter
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default
.PHONY: format

run: deps #TODO MD:swag-v1 proto-v1 ### swag run for API v1
	go mod download && \
	go run ./cmd/app
.PHONY: run

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
.PHONY: linter-hadolint

linter-dotenv: ### check by dotenv linter
	dotenv-linter
.PHONY: linter-dotenv

test: ### run test
	CGO_ENABLED=1 go test -v -race -covermode atomic -coverprofile=coverage.txt ./internal/...
.PHONY: test

#pkg_test: ### run pkg test
#	go clean -testcache && go test -v ./pkg/...

#integration-test: ### run integration-test
#	go clean -testcache && go test -v ./integration-test/...
#.PHONY: integration-test

bin-deps: ### install tools
	GOBIN=$(LOCAL_BIN) go install tool
.PHONY: bin-deps

pre-commit: format linter-golangci test ### run pre-commit
.PHONY: pre-commit
