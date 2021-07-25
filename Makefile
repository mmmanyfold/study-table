.PHONY: build clean deploy run push tag push format server

CONTAINER = study-table-service
COMMIT = $$(git rev-parse --short HEAD)
ECR_REPO_URI = 483649924232.dkr.ecr.us-east-1.amazonaws.com
IMAGE = $(CONTAINER):$(COMMIT)

build:
	@docker build \
#	--no-cache \
	--build-arg COMMIT=$(COMMIT) \
	--build-arg AIRTABLE_API_KEY=$$AIRTABLE_API_KEY \
	--build-arg AWS_ACCESS_KEY_ID=$$AWS_ACCESS_KEY_ID \
	--build-arg AWS_SECRET_ACCESS_KEY=$$AWS_SECRET_ACCESS_KEY \
	-t $(IMAGE) .

clean:
	echo "fix me"

deploy: clean build
	echo "fix me"

format:
	gofmt -w .

run:
	go install 
	study-table-service

tag:
	@docker tag study-table-service:$(COMMIT) $(ECR_REPO_URI)/study-table-service:latest

push:
	@docker push $(ECR_REPO_URI)/study-table-service:latest

server:
	@docker run -it -p 8080:8080 study-table-service:dev

