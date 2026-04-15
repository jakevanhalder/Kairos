package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	ActivePage string
}

func render(w http.ResponseWriter, page string, data PageData) {
	tmpl, err := template.ParseFiles(
		"internal/web/templates/base.html",
		"internal/web/templates/"+page+".html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/web/static"))))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		render(w, "index", PageData{ActivePage: "dashboard"})
	})
	mux.HandleFunc("GET /jobs", func(w http.ResponseWriter, r *http.Request) {
		render(w, "jobs", PageData{ActivePage: "jobs"})
	})
	mux.HandleFunc("GET /history", func(w http.ResponseWriter, r *http.Request) {
		render(w, "history", PageData{ActivePage: "history"})
	})
	mux.HandleFunc("GET /logs", func(w http.ResponseWriter, r *http.Request) {
		render(w, "logs", PageData{ActivePage: "logs"})
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
