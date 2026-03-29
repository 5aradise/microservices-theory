include .env
export $(shell sed 's/=.*//' .env)

export GOOSE_MIGRATION_DIR=sql/schema
export GOOSE_DBSTRING=host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable

export FRONT_END_BINARY=front

.PHONY: up down build_front start stop sqlc

## up: starts all containers in the background without forcing build
up: start
	@echo "Starting Docker images..."
	docker-compose up --build -d
	@echo "Docker images started!"

## down: stop docker compose
down: stop
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

sqlc:
	cd ./authentication && sqlc generate

migrate-up: 
	cd ./authentication && goose postgres up

migrate-down: 
	cd ./authentication && goose postgres down