package api

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph"
	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
