SHELL := /bin/bash

dev:
	@docker-compose -f deploy/docker-compose-dev.yml up -d && go run cmd/main.go

data:
	@docker-compose -f deploy/docker-compose-dev.yml up -d && go run scripts/populate.go

clean:
	@docker-compose -f deploy/docker-compose-dev.yml down; docker container prune -f && docker volume prune -f
