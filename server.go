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
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		urls, err := db.GetURLs(s.db)
		if err != nil {
			log.Fatalln(err)
		}

		urlsMap := make(map[string]string)

		for _, url := range urls {
			urlsMap[url.ID] = url.Link
		}

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
