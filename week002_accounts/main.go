package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type config struct {
	Pass string `env:"POSTGRES_PASSWORD"`
}

type server struct {
	conn *pgx.Conn
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(errors.Wrap(err, "error getting environment variables from file"))
	}

	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:pas@localhost:5432")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to database"))
	}
	defer conn.Close(context.Background())

	s := &server{
		conn: conn,
	}

	log.Println("running server on localhost:6969")
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
