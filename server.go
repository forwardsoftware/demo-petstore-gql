package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	dbqueries "petstore-gql/db"
	"petstore-gql/graph"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to DB
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Setup goose library
	goose.SetBaseFS(embedMigrations)

	// Setup goose dialect
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// Run migrations
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	// Setup DB queries
	queries := dbqueries.New(db)

	srv := handler.New(graph.NewExecutableSchema(graph.NewRootResolvers(queries)))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func connectDatabase() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("missing DATABASE_URL environment variable")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
