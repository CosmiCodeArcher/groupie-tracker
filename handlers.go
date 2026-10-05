package main

import (
	"groupie-tracker/api"
	"html/template"
	"log"
	"net/http"
)

type App struct {
	artists   []api.Artist
	relations []api.Relation
	tmpl      *template.Template
}

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "the page doesn't exist", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := a.tmpl.Execute(w, a.artists); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
