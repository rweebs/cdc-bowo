ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.PHONY: lint-prepare
lint-prepare:
	@echo "Preparing Linter"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINDIR=./.local/bin sh -s latest

.PHONY: lint
lint:
	@echo "Applying linter"
	./.local/bin/golangci-lint --timeout 5m run ./...

.PHONY: build
build:
	@go build -o data-schema-change

.PHONY: run
run:
	./data-schema-change run -c=config.test-app.json

# Short test is used for testing the whole unit-test
.PHONY: short-test
short-test:
	@go test -v -cover --short -race ./...

# Full test is used for testing the whole application including the database query directly to a live database
# This may takes time.
.PHONY: full-test
full-test:
	@echo "Running the full test..."
	@go test -v -cover -coverprofile coverage.out -race -p 1 ./...

run-all-tests:
	@go test -v -cover -coverprofile coverage.out -race -p 1 ./...