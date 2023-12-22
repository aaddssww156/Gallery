serve:
	@go run cmd/server/main.go

build:
	@go build -o bin/main cmd/server/main 

run:
	@./bin/main 