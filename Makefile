STORE_MANAGER_BOT_BINARY=storeManagerBotApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_store_manager_bot_backend
	@echo "Stopping docker images (if running...)"
	docker-compose down --remove-orphans
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_store_manager_bot_backend: builds the store manager backend binary as a linux executable
build_store_manager_bot_backend:
	@echo "Building store_manager_bot_backend binary..."
	cd backend && env GOOS=linux CGO_ENABLED=0 go build -o ${STORE_MANAGER_BOT_BINARY} ./cmd/api
	@echo "Done!"
