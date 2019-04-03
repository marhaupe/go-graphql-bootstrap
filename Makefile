output ?= binaries

domain ?= github.com
user ?= marhaupe
repository ?= go-graphql-bootstrap

dev:
	go run main.go

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

rename: 
	find . -type f | xargs perl -pi -e 's/github.com\/marhaup\
e\/go-graphql-bootstrap/$(domain)\/$(user)\/$(repository)/g;'
