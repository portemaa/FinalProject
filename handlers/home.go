package handlers

import (
	"net/http"
)

// Func serves as the main page
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/home.html")
}
