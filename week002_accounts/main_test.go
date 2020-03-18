package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	// external database should be prepared before running tests

	// create database for test purposes
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:pas@localhost:5432/postgres")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to database"))
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DROP DATABASE IF EXISTS testdb")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error removing test database"))
	}

	_, err = conn.Exec(context.Background(), "CREATE DATABASE testdb")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating test database"))
	}

	// run migration on prepared test database
	migr, err := migrate.New(
		"file://migrations",
		"postgres://postgres:pas@localhost:5432/testdb?sslmode=disable",
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error running migrations"))
	}

	err = migr.Up()
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot run migration"))
	}

	os.Exit(m.Run())
}

func migr() error {
	// run migrations

	return nil
}

//s := &server{
//conn: conn,
//}

//ts := httptest.NewServer(s.router())
//defer ts.Close()

//acc := &Account{
//Name:  "Neo",
//Email: "neo@the.one",
//}
//b, err := json.Marshal(acc)
//require.Nil(t, err)
//r := bytes.NewReader(b)
//client := &http.Client{}
//req, err := http.NewRequest("POST", ts.URL+"/acc", r)
//require.Nil(t, err)
//res, err := client.Do(req)
//require.Nil(t, err)
//acc2 := &Account{}

//dec := json.NewDecoder(res.Body)
//err = dec.Decode(acc2)
//require.Nil(t, err)

//fmt.Println("===>", acc2)
