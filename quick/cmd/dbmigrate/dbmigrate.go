package main

import (
	"app/gen/ent"
	"app/gen/ent/migrate"
	"app/quick"
	"context"
	"log"

	entsql "entgo.io/ent/dialect/sql"
)

func main() {
	plugins := quick.GetPlugins()

	drv := entsql.OpenDB(plugins.DB.Dialect, plugins.DB.DB)

	client := ent.NewClient(ent.Driver(drv), ent.Debug())

	defer client.Close()
	// Run the auto migration tool.
	err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false), // Disable foreign keys.
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
