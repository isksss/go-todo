APP_NAME = go-todo

all: tidy build

.PHONY: install
install:
	go install github.com/rakyll/statik@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:
	statik ./public
	sqlc generate --file sqlc/sqlc.yaml
	go build -o dist/${APP_NAME}
