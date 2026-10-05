package main

import (
	"groupie-tracker/api"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// stop the program on a fetch error so the server
	// does not serve error pages to a visitor
	artists, err := api.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}
	relationIndex, err := api.FetchRelations()
	if err != nil {
		log.Fatal(err)
	}

	app := &App{artists: artists, relations: relationIndex.Index, tmpl: tmpl}

	http.HandleFunc("/", app.home)
	log.Println("listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
