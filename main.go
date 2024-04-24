package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"text/template"
)

type Stock struct {
	Name  string
	Date  string
	Price string
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func dbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "stocks.db")
	if err != nil {
		log.Fatal(err.Error() + " During open stocks.db")
	}
	return db
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	rows, err := db.Query("SELECT * FROM STOCKS")
	if err != nil {
		log.Fatal(err.Error() + " Error with SELECT * FROM STOCKS")
	}

	defer rows.Close()

	stock := Stock{}
	res := []Stock{}

	for rows.Next() {
		var name, date, price string
		err := rows.Scan(&name, &date, &price)
		if err != nil {
			log.Fatal(err.Error())
		}

		stock.Name = name
		stock.Date = date
		stock.Price = price
		res = append(res, stock)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
}

func main() {
	log.Println("Server started on: http://localhost:80")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":80", nil)
}
