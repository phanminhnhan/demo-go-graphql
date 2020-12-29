package main

import (
	"graphql_postrgres/database"
	"graphql_postrgres/graph"
	"graphql_postrgres/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	//database
	sql := &database.Sql{
		Host:     "localhost",
		Port:      5432,
		Username: "postgres",
		Password: "postgres",
		DbName:   "graphql",
	}
	sql.ConnectDB()
	defer sql.Close()


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := generated.Config{Resolvers: &graph.Resolver{}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
