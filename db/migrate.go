package db

import (
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Use as a reference because it contains sync.WaitGroup
func (s *PostgresDBService) makeMigrations() {

	m, err := migrate.New(
		"file://db/migrations",
		s.connectionUrl)
	if err != nil {
		wlog.Fatalf("error creating migration object: %s", err.Error())
	}
	wlog.Infof("applying database migrations...")
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			wlog.Fatalf("error applying migrations: %s", err.Error())
		}
	}
	connErr, dbErr := m.Close()

	if connErr != nil {
		wlog.Fatalf("error closing migration connection: %s", connErr.Error())
	}
	if dbErr != nil {
		wlog.Fatalf("error closing migration database: %s", dbErr.Error())
	}

}
