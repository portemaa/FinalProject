package handlers

import (
	"html/template"
	"net/http"
	"password-checker/services"
)

func Generate(tmpl *template.Template, store services.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pwd := "example123!" // replace with generator logic
		store.SavePassword(pwd)

		data := map[string]interface{}{
			"Password": pwd,
		}
		// depends if we use html
		_ = tmpl.ExecuteTemplate(w, "generate.html", data)
	}
}
