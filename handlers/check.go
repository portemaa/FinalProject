package handlers

import (
	"net/http"
	"password-checker/services"
)

/*
REf: Acquired info from, Go Web Programming: Client-Server Model -> Handling Forms with net/http
*/

// Func handles the input of the password
// It reads the password from the request, checks the strength and then suggests improvements,
// completes the results into check.html tmeplate.

//var tmpl *template.Template

// Init sets the template from main.go, is a fix for earlier issues
/*func Init(t *template.Template) {
	tmpl = t
}*/

// Func handles POST and GET request
func Check(w http.ResponseWriter, r *http.Request) {
	// If the user submitted the form we POST request
	if r.Method == http.MethodPost {
		password := r.FormValue("password")                // Gets password value from the form
		strength := services.GetPasswordStrength(password) // Calls service side package and checks strength
		suggestions := services.ImprovePassword(password)  // Calls service side package and improves password

		// Puts results into a map for the template
		data := map[string]interface{}{
			"Password":    password,
			"Strength":    strength,
			"Suggestions": suggestions,
		}

		// Show the results in check.html
		if err := tmpl.ExecuteTemplate(w, "check.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If not POST, show the empty form
	if err := tmpl.ExecuteTemplate(w, "check.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
