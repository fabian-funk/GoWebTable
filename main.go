package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	db, _ := sql.Open("sqlite3", "stocks.db")
	rows, _ := db.Query("SELECT * FROM STOCKS")

	var stock, date, price string

	for rows.Next() {
		_ = rows.Scan(&stock, &date, &price)
		log.Printf("Result: %s: %s: %s", stock, date, price)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	log.Println("Server started on: http://localhost:80")
	http.ListenAndServe(":80", nil)
}
