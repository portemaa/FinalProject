package handlers

import (
	"html/template"
	"net/http"
	"password-checker/services"
	"strconv"
)

func Generate(tmpl *template.Template, store services.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pwd string

		if r.Method == http.MethodPost {
			// parse form and read length value
			if err := r.ParseForm(); err == nil {
				lengthStr := r.FormValue("length")
				length := 24 // default

				if lengthStr != "" {
					if n, err := strconv.Atoi(lengthStr); err == nil {
						//  min max bounds
						if n < 8 {
							n = 8
						} else if n > 32 {
							n = 32
						}
						length = n
					}
				}

				pwd = services.GeneratePassword(length)
				store.SavePassword(pwd)
			}
		}

		data := map[string]interface{}{
			"Password": pwd,
		}
		_ = tmpl.ExecuteTemplate(w, "generate.html", data)
	}
}
