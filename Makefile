.PHONY: build
build:
	@go env -w CGO_ENABLED=0 
	@go build -ldflags="-w -s" -o main