package foo

import (
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/vibec0dr/snippetbox/internal/config"
)

func ExampleHandler(app config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go")

		files := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			"./ui/html/pages/home.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			app.Logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
