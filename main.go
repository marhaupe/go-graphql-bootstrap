package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/marhaupe/go-graphql-bootstrap/server"
)

var port = os.Getenv("PORT")
var defaultPort = "3000"

func main() {
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", handler.Playground("GraphQL playground", "/query"))
	http.HandleFunc("/query", server.QueryHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
