package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matryer/is"
	"github.com/yevhenshymotiuk/bluefield-url-shortener-task/db"
)

func newDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}

	return db, mock
}

func TestHandleIndex(t *testing.T) {
	is := is.New(t)
	db, _ := newDBMock()

	srv, err := newServer(db)
	is.NoErr(err)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	srv.router.ServeHTTP(w, req)

	is.Equal(w.Result().StatusCode, http.StatusOK)
}

func TestHandleShortenedURL(t *testing.T) {
	is := is.New(t)
	database, mock := newDBMock()
	defer database.Close()

	srv, err := newServer(database)
	is.NoErr(err)

	url := db.URL{ID: db.NewURLID(), Link: "http://example.com/"}

	addURLQuery := regexp.QuoteMeta("INSERT INTO URL(ID, link) values (?, ?)")
	getURLQuery := regexp.QuoteMeta("SELECT * FROM URL WHERE ID=?")
	getURLQueryRows := sqlmock.NewRows(
		[]string{"ID", "link"},
	).AddRow(
		url.ID,
		url.Link,
	)

	_ = sqlmock.NewRows([]string{"ID", "link"}).AddRow(url.ID, url.Link)

	mock.ExpectPrepare(addURLQuery)
	mock.ExpectExec(
		addURLQuery,
	).WithArgs(
		url.ID,
		url.Link,
	).WillReturnResult(
		sqlmock.NewResult(1, 1),
	)

	err = db.AddURL(database, url)
	is.NoErr(err)

	mock.ExpectQuery(
		getURLQuery,
	).WithArgs(
		url.ID,
	).WillReturnRows(
		getURLQueryRows,
	)

	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", url.ID), nil)
	w := httptest.NewRecorder()

	srv.router.ServeHTTP(w, req)

	is.Equal(w.Result().StatusCode, http.StatusFound)
}
