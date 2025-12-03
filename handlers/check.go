package handlers

import (
	"html/template"
	"net/http"
	"password-checker/services"
)

/*
REf: Acquired info from, Go Web Programming: Client-Server Model -> Handling Forms with net/http
*/

// Func handles the input of the password
// It reads the password from the request, checks the strength and then suggests improvements,
// completes the results into check.html tmeplate.
func Check(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Checks if the request is a POST method( or submitted on the template)
		if r.Method == http.MethodPost {
			password := r.FormValue("password")

			// Calls service layer: checks analyze strength and improvements
			strength := services.GetPasswordStrength(password)
			suggestions := services.ImprovePassword(password)

			// Creates a map and send data into the HTML template.
			data := map[string]interface{}{
				"Password":    password,
				"Strength":    strength,
				"Suggestions": suggestions,
			}

			// Tells project to fill in the check.html template
			if err := tmpl.ExecuteTemplate(w, "check.html", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		// If GET request, template shows the empty form
		// This code runs only if the request method is not POST Method
		if err := tmpl.ExecuteTemplate(w, "check.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
