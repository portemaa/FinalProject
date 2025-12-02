package handlers

import (
	"html/template"
	"net/http"
)

func Generate(tmpl *template.Template, store *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := services.GeneratePassword(16)
		store.SavePassword(password)

		data := map[string]interface{}{
			"Generated": password,
		}
		// depends if we use html
		_ = tmpl.ExecuteTemplate(w, "generate.html", data)
	}
}
