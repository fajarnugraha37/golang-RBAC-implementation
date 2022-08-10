install:
	go mod tidy
build:
	go build main.go
dev:
	go run main.go