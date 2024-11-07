.PHONY:
.DEFAULT_GOAL := build

build:
	go env -w CGO_ENABLED=0
	go env -w GOOS=linux
	go mod download && go build -o ./.bin/reminder-app ./cmd/main.go

run: build
	docker-compose up --remove-orphans reminder-app 
