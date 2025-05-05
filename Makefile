# Set Variables
APP_NAME=notes-api
ACR_NAME=notes
ACR_URL=$(ACR_NAME).azurecr.io
DOCKER_IMAGE=$(ACR_URL)/$(APP_NAME)

# Default Target
all: build run

# Build Go Project
build:
    @echo "Building $(APP_NAME)..."
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

# Build Docker Image
docker-build:
    @echo "Building Docker image..."
    docker build -t $(DOCKER_IMAGE):latest .

# Push Docker Image to Azure Container Registry (ACR)
docker-push:
    @echo "Logging in to Azure Container Registry..."
	az acr login --name $(ACR_NAME)
    @echo "Pushing Docker image to ACR..."
    docker push $(DOCKER_IMAGE):latest
    @echo "Tagging image with commit SHA..."
    docker tag $(DOCKER_IMAGE):latest $(DOCKER_IMAGE):$(shell git rev-parse HEAD)
    docker push $(DOCKER_IMAGE):$(shell git rev-parse HEAD)

# Clean Build Artifacts
clean:
    @echo "Cleaning build artifacts..."
    rm -f $(APP_NAME)
