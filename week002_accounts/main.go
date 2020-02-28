package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type server struct {
	conn *pgx.Conn
}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to database"))
	}
	defer conn.Close(context.Background())

	s := &server{
		conn: conn,
	}

	log.Fatal(http.ListenAndServe(":6969", s.router()))
}

func sendErr(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}

func (s *server) router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.handleHome())

	return r
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there!"))
	}
}

type Account struct {
	Name  string
	Email string
	Pass  string
}

func (s *server) handleAccountCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acc := &Account{}

		err := json.NewDecoder(r.Body).Decode(acc)
		if err != nil {
			sendErr(w, "invalid request body")
			return
		}

	}
}
