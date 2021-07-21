.PHONY: build clean deploy run push tag push docker-build

CONTAINER = study-table-service
COMMIT = $$(git rev-parse --short HEAD)
ECS_REPO_URI = 483649924232.dkr.ecr.us-east-1.amazonaws.com
IMAGE = $(CONTAINER):dev

build:
	go install

docker-build: 
	@docker build \
	--no-cache \
	--build-arg COMMIT=$(COMMIT) \
	--build-arg AIRTABLE_API_KEY=$$AIRTABLE_API_KEY \
	-t $(IMAGE) .

clean:
	echo "fix me"

deploy: clean build
	echo "fix me"

format:
	go fmt .

run:
	go install 
	study-table-service

tag:
	echo "fix me"

push:	
	echo "fix me"
