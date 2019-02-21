output ?= server

build:
	go build -o $(output)

regen:
	go run scripts/gqlgen.go

test:
	go test ./...

coverage:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html 
	rm cover.out
