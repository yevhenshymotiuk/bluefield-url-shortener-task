package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yevhenshymotiuk/bluefield-url-shortener-task/db"
)

type server struct {
	db     *sql.DB
	router *httprouter.Router
}

func newServer(db *sql.DB) (*server, error) {
	router := httprouter.New()

	s := &server{db: db, router: router}
	s.routes()

	return s, nil
}

func (s *server) handleIndex() httprouter.Handle {
	type data struct {
		Host string
		ID   string
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		longURL := r.Form.Get("url")

		var shortenedURL db.URL

		if longURL != "" {
			shortenedURL, err = db.AddURL(s.db, db.URL{Link: longURL})
			if err != nil {
				log.Fatalln(err)
			}
		}

		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			log.Fatalln(err)
		}

		t.Execute(w, data{Host: r.Host, ID: shortenedURL.ID})
	}
}

func (s *server) handleShortenedURL() httprouter.Handle {
	var emptyURL db.URL

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")

		url, err := db.GetURL(s.db, id)
		if err != nil {
			log.Fatalln(err)
		}

		if url == emptyURL {
			http.NotFound(w, r)
		} else {
			http.Redirect(w, r, url.Link, http.StatusFound)
		}
	}
}
