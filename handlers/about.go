package handlers

import (
	"net/http"
)

// About renders the about page using the shared `tmpl` initialized by `handlers.Init`.
// this is simpler than the previous code
func About(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "about.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
