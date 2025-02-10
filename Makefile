SHELL := /bin/bash

dev:
	@docker-compose -f deploy/docker-compose-dev.yml up -d && sleep 3 && go run cmd/main.go

users:
	@go run scripts/populate_db/users.go

clean:
	@docker-compose -f deploy/docker-compose-dev.yml down; docker container prune -f && docker volume prune -f
