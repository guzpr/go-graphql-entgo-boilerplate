package gql

import (
	"github.com/99designs/gqlgen/graphql"
	ur "github.com/sekalahita/epirus/internal/domain/appuser/gql"
	"github.com/sekalahita/epirus/internal/ent/gen"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client       *gen.Client
	userResolver ur.Resolver
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *gen.Client) graphql.ExecutableSchema {
	c := Config{
		Resolvers: &Resolver{
			client:       client,
			userResolver: ur.NewResolver(client),
		},
	}
	return NewExecutableSchema(c)
}
