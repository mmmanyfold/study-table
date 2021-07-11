.PHONY: build clean deploy

build: 
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main src/main.go

clean:
	echo "fix me"

deploy: clean build
	echo "fix me"

format:
	go fmt ./src