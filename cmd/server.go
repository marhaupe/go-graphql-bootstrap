package cmd

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/marhaupe/go-graphql-bootstrap/internal"
)

const defaultPort = "8080"

func Run(port string) {
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(internal.NewExecutableSchema(internal.Config{Resolvers: &internal.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
