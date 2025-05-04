# Set Variables
APP_NAME=notes-api
DOCKER_IMAGE=notes.azurecr.io/$(APP_NAME)

# Default Target
all: build run

# Build Go Project
build:
	@echo $(APP_NAME)
	go build -o $(APP_NAME)

# Run Application
run:
	./$(APP_NAME)

# Format Code
fmt:
	go fmt ./...

# Run Tests
test:
	go test ./...

# Build & Push Docker Image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-push:
	az acr login --name notes
	docker push $(DOCKER_IMAGE)

# Deploy to AKS
deploy:
	kubectl set image deployment/$(APP_NAME) $(APP_NAME)=$(DOCKER_IMAGE):latest

# Clean Build Artifacts
clean:
	rm -f $(APP_NAME)
