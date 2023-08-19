package migrate

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateUp(postgresDB *sql.DB, path string, dbName string, migrationsTable string, steps *int64) error {
	migrator, err := newMigrator(postgresDB, path, dbName, migrationsTable)
	if err != nil {
		return err
	}

	if steps != nil {
		if err := migrator.Steps(int(*steps)); !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		return nil
	}

	if err := migrator.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func MigrateDown(postgresDB *sql.DB, path string, dbName string, migrationsTable string) error {
	migrator, err := newMigrator(postgresDB, path, dbName, migrationsTable)
	if err != nil {
		return err
	}

	if err := migrator.Down(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func newMigrator(db *sql.DB, path string, dbName string, migrationsTable string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: migrationsTable,
	})
	if err != nil {
		return nil, err
	}
	if !strings.Contains(path, "file://") {
		path = "file://" + path
	}

	return migrate.NewWithDatabaseInstance(path, dbName, driver)
}
