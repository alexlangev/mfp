.DEFAULT_GOAL := build

.PHONY:fmt vet build clean test
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o mfp ./cmd/mfp/

clean:
	go clean
	rm ./mfp

test:
	go test -v ./...
