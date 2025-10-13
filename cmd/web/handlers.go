package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		app.serverError(w, r, err)
	}

}

func (app application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID: %d\n", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Display a form for creating a new snippet...")
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Save a new snippet")
}
