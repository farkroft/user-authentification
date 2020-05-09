NAME := auth-service
TAG := $$(git log -1 --pretty=%H(MISSING))
IMG := ${NAME}:${TAG}
LATEST := ${NAME}:latest

run:
	docker-compose up

test:
	go test ./... -cover