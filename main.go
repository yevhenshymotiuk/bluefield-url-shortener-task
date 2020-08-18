package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yevhenshymotiuk/bluefield-url-shortener-task/db"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		return err
	}
	defer database.Close()

	if _, err = os.Stat("./db/db.sqlite3"); os.IsNotExist(err) {
		err = db.Init(database)
		if err != nil {
			return err
		}
	}

	urls, err := db.GetURLs(database)
	if err != nil {
		return err
	}
	fmt.Println(urls)

	srv := newServer()

	port := 8080
	log.Println(fmt.Sprintf("Listening on %d", port))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), srv.router)
}
