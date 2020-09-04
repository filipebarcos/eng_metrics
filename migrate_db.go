package eng_metrics

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	tern "github.com/jackc/tern/migrate"
	"github.com/stoplightio/eng_metrics/db"
)

// MigrateDb
func MigrateDb(w http.ResponseWriter, _r *http.Request) {
	log.Println("Starting DB schema migration")
	ctx := context.Background()

	conn, err := db.NewConn()
	if err != nil {
		log.Printf("An error occurred trying to get a new DB connection: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close(ctx)

	log.Println("DB connection established")

	migrator, err := tern.NewMigrator(ctx, conn, "public.schema_version")
	if err != nil {
		log.Printf("An error occurred when getting a new Migrator instance: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Migrator created")

	migrator.LoadMigrations(filepath.Join(os.Getenv("SL_CODEBASE_ROOT"), "migrate"))

	if len(migrator.Migrations) == 0 {
		log.Println("No migrations found")
		http.Error(w, "no migraitons found", http.StatusInternalServerError)
		return
	} else {
		log.Println("Migrations loaded")
	}

	migrator.OnStart = func(sequence int32, name string, direction string, sql string) {
		log.Printf("Running migration: %s (%s)\n", name, direction)
	}

	log.Println("Will call Migrate")
	err = migrator.Migrate(ctx)

	if err != nil {
		log.Printf("An error occurred when running migrations: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Done running migrations")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"ok": true}`)
}
