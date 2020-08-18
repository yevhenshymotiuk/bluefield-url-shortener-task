// Package db provides functions and data sctructures
// to operate on a database
package db

import (
	"database/sql"
	"log"
	"strings"

	"github.com/google/uuid"
)

type URL struct {
	ID   string
	Link string
}

// Init initializes a database
func Init(db *sql.DB) error {
	log.Println("DB was initialized")

	err := createURLTable(db)
	if err != nil {
		return err
	}

	err = AddURL(db, URL{Link: "https://pkg.go.dev/"})

	return err
}

func createURLTable(db *sql.DB) error {
	createURLTableSQL := `CREATE TABLE URL (
    "ID" UUID NOT NULL PRIMARY KEY,
    "link" TEXT
);`

	statement, err := db.Prepare(createURLTableSQL)
	if err != nil {
		return err
	}
	statement.Exec()

	return nil
}

// GetURLs returns all URLs
func GetURLs(db *sql.DB) ([]URL, error) {
	var URLs []URL

	row, err := db.Query("SELECT * FROM URL ORDER BY link")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		var (
			ID   string
			link string
		)

		row.Scan(&ID, &link)

		URLs = append(URLs, URL{ID: ID, Link: link})
	}

	return URLs, nil
}

// AddURL adds new URL
func AddURL(db *sql.DB, url URL) error {
	insertURLSQL := `INSERT INTO URL(ID, link) values (?, ?)`

	statement, err := db.Prepare(insertURLSQL)
	if err != nil {
		return err
	}

	id := strings.Replace(uuid.New().String(), "-", "", -1)[:4]

	_, err = statement.Exec(id, url.Link)

	return err
}
