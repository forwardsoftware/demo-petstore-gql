package graph

import (
	"petstore-gql/db"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	q *db.Queries
}

func NewRootResolvers(q *db.Queries) Config {
	return Config{
		Resolvers: &Resolver{
			q: q,
		},
	}
}
