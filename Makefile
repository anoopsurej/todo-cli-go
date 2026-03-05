.PHONY: check build test clean

check:
	go vet ./...

build:
	go build -o tasks ./cmd/tasks/

test:
	go test ./...

clean:
	rm -f tasks