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
var db, _ = sql.Open("sqlite3", "stocks.db")

func Index(w http.ResponseWriter, r *http.Request) {

	selDB, err := db.Query("SELECT * FROM STOCKS")
	if err != nil {
		panic(err.Error())
	}

	stock := Stock{}
	res := []Stock{}

	for selDB.Next() {
		var name, date, price string
		err := selDB.Scan(&name, &date, &price)
		if err != nil {
			panic(err.Error())
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
