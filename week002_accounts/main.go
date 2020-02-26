package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct{}

func main() {
	s := &Server{}

	log.Fatal(http.ListenAndServe(":6969", s.router()))
}

func (s *Server) router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.handleHome())

	return r
}

func (s *Server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there!"))
	}
}
