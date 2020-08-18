package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	router *httprouter.Router
}

func newServer() *server {
	router := httprouter.New()

	s := &server{router: router}
	s.routes()

	return s
}

func (s *server) handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Hello, World!")
	}
}
