package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yevhenshymotiuk/bluefield-url-shortener-task/db"
)

type server struct {
	db *sql.DB
	router *httprouter.Router
}

func newServer() (*server, error) {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		return nil, err
	}
	defer database.Close()

	if _, err = os.Stat("./db/db.sqlite3"); os.IsNotExist(err) {
		err = db.Init(database)
		if err != nil {
			return nil, err
		}
	}

	router := httprouter.New()

	s := &server{db: database, router: router}
	s.routes()

	return s, nil
}

func (s *server) handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Hello, World!")
	}
}

func (s *server) handleShortenedURL() httprouter.Handle {
	urls, err := db.GetURLs(s.db)
	if err != nil {
		log.Fatalln(err)
	}

	urlsMap := make(map[string]string)

	for _, url := range urls {
		urlsMap[url.ID] = url.Link
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		if err != nil {
			log.Fatalln(err)
		}

		url, prs := urlsMap[id]
		if prs {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			http.NotFound(w, r)
		}
	}
}
