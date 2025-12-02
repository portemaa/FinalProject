package handlers

import (
	"html/template"
	"net/http"
)

func Generate(tmpl *template.Template, store *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := logic.GeneratePassword(16)
		store.SavePassword(password)

		data := map[string]interface{}{
			"Generated": password,
		}

		_ = tmpl.ExecuteTemplate(w, "generate.html", data)
	}
}
