package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	databaseURL := "postgres://ys-user:qwerty@localhost:5432/ys-db?sslmode=disable"
	m, err := migrate.New("file://tools/migrate/migrations", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// m.Down() - to discard changes
	// m.Force() - to fix dirty version of migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied migration: %d, Dirty: %t\n", version, dirty)
}
