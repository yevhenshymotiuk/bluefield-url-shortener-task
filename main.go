package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yevhenshymotiuk/bluefield-url-shortener-task/db"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	database, err := db.Setup()
	if err != nil {
		return err
	}
	defer database.Close()

	srv, err := newServer(database)
	if err != nil {
		return err
	}
	defer srv.db.Close()

	port := 8080
	log.Println(fmt.Sprintf("Listening on %d", port))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), srv.router)
}
