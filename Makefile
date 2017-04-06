all: lint test

build:
	gb build

vendor:
	gb vendor restore

test:
	gb test -v

lint:
	go vet ./src/...
	golint