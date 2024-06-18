package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		fmt.Println("POSTGRES_URL must be set")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		row := db.QueryRow("SELECT 1")

		var one int
		err := row.Scan(&one)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		fmt.Fprintf(w, "OK")
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
