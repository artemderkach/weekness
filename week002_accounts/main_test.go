package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var testDatabase = "testdb"

func TestMain(t *testing.T) {
	require.Nil(t, initDB(testDatabase))
	defer func() {
		require.Nil(t, dropDB())
	}()

	url := fmt.Sprintf("postgresql://postgres:pas@localhost:5432/%s", testDatabase)
	conn, err := pgx.Connect(context.Background(), url)
	require.Nil(t, errors.Wrap(err, "error connecting to db"))
	defer func() {
		require.Nil(t, conn.Close(context.Background()))
	}()

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
func initDB(dbname string) error {
	// create database for test purposes
	// firstly connect to default one, than create test db
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

	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return errors.Wrap(err, "error removing test database")
	}

	query = fmt.Sprintf("CREATE DATABASE %s", dbname)
	fmt.Println("---->", query)
	ct, err := conn.Exec(context.Background(), query)
	fmt.Println("====>", ct)
	if err != nil {
		return errors.Wrap(err, "error creating test database")
	}

	// run migration on prepared test database
	dbURL := fmt.Sprintf("postgres://postgres:pas@localhost:5432/%s?sslmode=disable", dbname)
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

	// write test data into database
	url := fmt.Sprintf("postgresql://postgres:pas@localhost:5432/%s", testDatabase)
	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		return errors.Wrapf(err, "cannot connect to test database '%s'", dbname)
	}
	defer func() {
		err = conn.Close(context.Background())
		if err != nil {
			log.Println(errors.Wrap(err, "error closing connection"))
		}
	}()

	file, err := os.Open("./migrations/testdata/users.sql")
	if err != nil {
		return errors.Wrap(err, "error opening dump file")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(errors.Wrap(err, "error closing file"))
		}
	}()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(byte(';'))
		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			return errors.Wrap(err, "error parsing bump file")
		}

		_, err = conn.Exec(context.Background(), line)
		if err != nil {
			return errors.Wrap(err, "cannot execute query from dump file")
		}

	}

	return nil
}

func dropDB() error {
	url := fmt.Sprintf("postgresql://postgres:pas@localhost:5432/postgres")
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return errors.Wrap(err, "error connecting to database")
	}

	defer func() {
		err = conn.Close(context.Background())
		if err != nil {
			log.Println(errors.Wrap(err, "error closing connection"))
		}
	}()

	query := fmt.Sprintf("DROP DATABASE %s", testDatabase)
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return errors.Wrap(err, "error dropping test database")
	}

	return nil
}
