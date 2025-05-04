package main

import (
	"flag"
	"log"

	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var (
		sourcePath  string
		databaseUrl string
	)
	flag.StringVar(&sourcePath, "source-path", "", "source path")
	flag.StringVar(&databaseUrl, "database-url", "", "database url")
	flag.Parse()
	if sourcePath == "" {
		log.Fatal("source-path is required")
	}
	if databaseUrl == "" {
		log.Fatal("database-url is required")
	}

	m, err := migrate.New(
		"file://"+sourcePath,
		databaseUrl)
	if err != nil {

		log.Fatal(errors.Wrap(err, "migration failed"))
	}
	if err := m.Up(); err != nil {
		log.Fatal(errors.Wrap(err, "migration failed"))
	}
}
