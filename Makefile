APP_NAME=jwt-auth

build:
	@go build -o bin/$(APP_NAME) cmd/api/main.go

run: build
	@./bin/$(APP_NAME)

dev:
	go run cmd/api/main.go

test:
	@go test -v ./...

migrate:
	go run cmd/migrate/main.go