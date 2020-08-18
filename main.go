package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	srv := newServer()

	port := 8080
	log.Println(fmt.Sprintf("Listening on %d", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), srv.router)
}
