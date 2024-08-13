package main

import (
	"log"

	"github.com/iancoleman/strcase"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var err error

	opts := []entc.Option{
		// entc.TemplateDir("./cmd/entg/tmpl"),
		entc.FeatureNames("sql/execquery", "schema/snapshot", "sql/modifier", "sql/upsert"),
	}

	err = entc.Generate("../schema",
		&gen.Config{
			Target:  "../gen/ent",
			Package: "app/gen/ent",
			Hooks: []gen.Hook{
				FieldSnakeToCamel(),
			},
		},
		opts...,
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

// convert ent snake_case json tags to camelCase for all the schema fields automatically
// https://github.com/ent/ent/issues/1069
func FieldSnakeToCamel() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					field.StructTag = `json:"` + strcase.ToLowerCamel(field.Name) + `,omitempty"`
				}
			}
			return next.Generate(g)
		})
	}
}
