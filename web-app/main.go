package main

import (
	"log"
	"net/http"
	"text/template"
)

type Products struct {
	Name        string
	Description string
	Price       int
	Qty         int
}

var temp = template.Must(template.ParseGlob("public/*.html"))

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	products := []Products{
		{"T-Shirt", "Pretty", 29, 10},
		{"Laptop", "Lenovo", 1000, 5},
		{"Mouse", "Razer", 100, 0},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
