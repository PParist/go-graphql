package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PParist/go-graphql/graph"
	"github.com/labstack/echo/v4"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	queryHandler := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{}},
	))

	e.POST("/query", func(c echo.Context) error {
		queryHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", queryHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(e.Start(":" + port))

}
