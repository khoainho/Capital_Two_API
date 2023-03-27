build:
	@go build -o bin/Capital_Two

run: build
	@./bin/Capital_Two

test:
	@go test -v ./...