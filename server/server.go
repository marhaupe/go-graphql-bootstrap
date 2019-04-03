package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"

	"github.com/99designs/gqlgen/handler"
)

func QueryHandler() http.HandlerFunc {
	schema := NewExecutableSchema(Config{Resolvers: &rootResolver{}})

	withLogging := handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
		nextRes := next(ctx)
		requestContext := graphql.GetRequestContext(ctx)

		query := requestContext.RawQuery

		// Only log the query if it's not an introspect query sent by playgrounds
		if !strings.HasPrefix(query, "query IntrospectionQuery") {
			log.Printf("Query: %v", query)
			log.Printf("Variables: %+v", requestContext.Variables)
			log.Printf("Result: %v", string(nextRes))
		}

		return nextRes
	})

	return handler.GraphQL(schema, withLogging)
}
