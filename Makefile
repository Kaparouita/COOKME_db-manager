server:
	go run main.go

build-db:
	go build -o bin/server main.go

include ./image.mk