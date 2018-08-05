.PHONY: dep build run

dep:
	dep ensure

build:
	go build

run:
	go run main.go
