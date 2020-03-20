package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var testDatabase = "testdb"

func TestMain(t *testing.T) {
	err := initDB()
	require.Nil(t, err)
	err = dropDB()
	require.Nil(t, err)

	url := fmt.Sprintf("postgresql://postgres:pas@localhost:5432/%s", testDatabase)
	conn, err := pgx.Connect(context.Background(), url)
	require.Nil(t, errors.Wrap(err, "error connecting to db"))
	defer conn.Close(context.Background())

	srv := &server{
		conn: conn,
	}

	ts := httptest.NewServer(srv.router())

	acc := &Account{
		Name:  "Neo",
		Email: "TheOne@gmail.com",
	}

	accBody, err := json.Marshal(acc)
	require.Nil(t, err)

	r := bytes.NewReader(accBody)
	res, err := http.Post(ts.URL+"/acc", "", r)
	require.Nil(t, err)

	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	require.Equal(t, "created", string(body))
}

// initDB prepares database for running tests
// - create test database
// - run migrations
// - run sql dump (test data)
func initDB() error {
	// create database for test purposes
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:pas@localhost:5432/postgres")
	if err != nil {
		return errors.Wrap(err, "error connecting to database")
	}

	defer func() {
		err = conn.Close(context.Background())
		if err != nil {
			log.Println(errors.Wrap(err, "error closing connection"))
		}
	}()

	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return errors.Wrap(err, "error removing test database")
	}

	query = fmt.Sprintf("CREATE DATABASE %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return errors.Wrap(err, "error creating test database")
	}

	// run migration on prepared test database
	dbURL := fmt.Sprintf("postgres://postgres:pas@localhost:5432/%s?sslmode=disable", testDatabase)
	migr, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		return errors.Wrap(err, "error running migrations")
	}

	defer func() {
		srcErr, dbErr := migr.Close()
		if srcErr != nil {
			log.Println(errors.Wrap(srcErr, "error closing connection"))
		}
		if dbErr != nil {
			log.Println(errors.Wrap(dbErr, "error closing connection"))
		}
	}()

	err = migr.Up()
	if err != nil {
		return errors.Wrap(err, "cannot run migration")
	}

	return nil
}

func dropDB() error {
	url := fmt.Sprintf("postgresql://postgres:pas@localhost:5432/%s", testDatabase)
	conn, err := pgx.Connect(context.Background())
	if err != nil {
		return errors.Wrap(err, "error connecting to database")
	}

	defer func() {
		err = conn.Close(context.Background())
		if err != nil {
			log.Println(errors.Wrap(err, "error closing connection"))
		}
	}()

	query = fmt.Sprintf("DROP DATABASE %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return errors.Wrap(err, "error dropping test database")
	}

	return nil
}
