.PHONY: build

build:
	docker-compose build front

.PHONY: status logs start stop clean

ps:
	docker-compose ps

logs:
	docker-compose logs -f front

up:
	docker-compose up -d

start:
	docker-compose start front

stop:
	docker-compose stop front

down:stop
	docker-compose down -v --remove-orphans

attach:
	docker-compose exec front bash

prune:
	docker system prune --all --volumes

.PHONY: gtest

gtest:
	go test -v -cover -coverprofile coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
