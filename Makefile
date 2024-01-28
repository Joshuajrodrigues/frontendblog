build:
	go build -o bin/blog

run: build
	./bin/blog

test:
	go test -v ./...