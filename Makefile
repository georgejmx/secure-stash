build:
	go build -o bin/secure-stash .

test:
	go test ./...

run-source: test
	go run main.go

run: build
	./bin/secure-stash
