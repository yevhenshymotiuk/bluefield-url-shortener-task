// Package db provides functions and data sctructures
// to operate on a database
package db

import (
	"database/sql"
	"log"
)

type URL struct {
	ID   int
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
	createURLTableSQL := `CREATE TABLE IF NOT EXISTS URL (
    "ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
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
			ID   int
			link string
		)

		row.Scan(&ID, &link)

		URLs = append(URLs, URL{ID: ID, Link: link})
	}

	return URLs, nil
}

// AddURL adds new URL
func AddURL(db *sql.DB, url URL) error {
	insertURLSQL := `INSERT INTO URL(link) values (?)`

	statement, err := db.Prepare(insertURLSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(url.Link)

	return err
}
