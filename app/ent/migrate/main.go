//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/Shota-Kurahashi/ent-with-echo/ent/migrate"
	"github.com/Shota-Kurahashi/ent-with-echo/src/config"
	_ "github.com/go-sql-driver/mysql"
)

// func main() {
// 	client := db.Connect()
// 	defer client.Close()

// 	ctx := context.Background()
// 	// マイグレーションの実行
// 	err := client.Schema.Create(
// 		ctx,
// 		migrate.WithDropIndex(true),
// 		migrate.WithDropColumn(true),
// 	)
// 	if err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// }

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, config.GetAtlasUrl(), os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
