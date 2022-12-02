package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("public/*.html"))

func dbConnect() *sql.DB {
	connection := "user=postgres dbname=postgres password=postgres host=host.docker.internal sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Products struct {
	Name        string
	Description string
	Price       float64
	Ammount     int
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	rows, err := db.Query("select * from products")
	if err != nil {
		log.Fatal(err)
	}
	p := Products{}
	products := []Products{}

	for rows.Next() {
		var id, ammount int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &ammount)
		if err != nil {
			log.Fatal(err)
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Ammount = ammount

		products = append(products, p)
	}
	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
