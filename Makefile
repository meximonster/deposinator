SHELL := /bin/bash

run:
	@docker-compose -f deploy/docker-compose-dev.yml up -d && go run cmd/main.go

users:
	@go run scripts/populate_db/users.go
