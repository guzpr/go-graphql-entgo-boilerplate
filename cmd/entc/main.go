//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../../internal/gql/ent.graphqls"),
		entgql.WithConfigPath("../../internal/gql/gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("entgql.NewExtension(): %v", err)
	}

	err = entc.Generate(
		"../../internal/ent/schema",
		&gen.Config{
			Target:  "../../internal/ent/gen",
			Package: "github.com/sekalahita/epirus/internal/ent/gen",
			Features: []gen.Feature{
				gen.FeatureModifier,
				gen.FeatureIntercept,
				// gen.FeatureSnapshot,
			},
		},
		entc.Extensions(ex),
	)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
