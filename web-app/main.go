package main

import (
	"log"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("public/*.html"))

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
