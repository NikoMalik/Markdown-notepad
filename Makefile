build:
	@go build -o bin/markdown

run: build
	@./bin/markdown

test:
	@go test -v ./...