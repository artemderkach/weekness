package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

var testDatabase = "testdb"

// TestMain prepares database for running tests
// - create database
// - run migrations
// - run tests
// - close connections and drop database
func TestMain(m *testing.M) {
	// create database for test purposes
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:pas@localhost:5432/postgres")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to database"))
	}

	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error removing test database"))
	}

	query = fmt.Sprintf("CREATE DATABASE %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating test database"))
	}

	// run migration on prepared test database
	dbURL := fmt.Sprintf("postgres://postgres:pas@localhost:5432/%s?sslmode=disable", testDatabase)
	migr, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error running migrations"))
	}

	err = migr.Up()
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot run migration"))
	}

	code := m.Run()

	// clean up connections and db
	srcErr, dbErr := migr.Close()
	if srcErr != nil {
		log.Fatal(errors.Wrap(srcErr, "error closing connection"))
	}
	if dbErr != nil {
		log.Fatal(errors.Wrap(dbErr, "error closing connection"))
	}

	query = fmt.Sprintf("DROP DATABASE %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error dropping test database"))
	}
	err = conn.Close(context.Background())
	if err != nil {
		log.Fatal(errors.Wrap(err, "error closing connection"))
	}

	defer os.Exit(code)
}
