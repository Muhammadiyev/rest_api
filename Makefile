.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

migrate-up:
	migrate -path ./migrations -database 'postgres://rest_api:qwerty1234!@localhost:5432/restapi_dev?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://rest_api:qwerty1234!@localhost:5432/restapi_dev?sslmode=disable' down