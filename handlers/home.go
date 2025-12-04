package handlers

import "net/http"

// Func is the handler for the root URL ("/")
// (Version 2) home.html renders through the same template as the other pages,
//so {{template "navbar"}} (added to all templates) works consistently everywhere

func Home(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
