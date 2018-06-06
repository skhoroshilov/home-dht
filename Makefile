BINARY_NAME=home-dht

all: clean build check test

clean:
	go clean
	rm -f $(BINARY_NAME)

build:
	go build -o $(BINARY_NAME)

test:
	go test -v ./...

dep:
	go test -i
	go get -u ./...
	go get -u github.com/kisielk/errcheck

check:
	go vet ./...
	errcheck ./...
