package main

import (
	"github.com/go-pg/pg/v9"
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


	DB := database.New(&pg.Options{
		User: "postgres",
		Password: "postgres",
		Database: "graphql",
	})
	defer DB.Close()

	DB.AddQueryHook(database.DBLogger{})


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := generated.Config{Resolvers: &graph.Resolver{
		LinkRepo: database.LinkRepo{Db: DB},
		UserRepo: database.UserRepo{Db: DB},
	}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
