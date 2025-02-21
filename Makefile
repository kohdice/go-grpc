BIN := ./bin/server
CMD := ./cmd/server/main.go

.PHONY: all
all: clean build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BIN) $(CMD)

.PHONY: clean
clean:
	go clean

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	goimports -w .
	go fmt ./...

.PHONY: lint
lint:
	go vet ./...

.PHONY: ci-lint
ci-lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: coverage
coverage:
	go test -v -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

