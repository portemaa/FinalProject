package handlers

import (
	"net/http"
	"password-checker/services"
)

var store services.StorageService

// Func sets the storage from main.go
func InitStore(s services.StorageService) {
	store = s
}

// Func shows all stored passwords and basic counts
func Stats(w http.ResponseWriter, r *http.Request) {
	// Get all passwords from storage
	passwords := store.GetAll()

	data := map[string]interface{}{
		"Passwords":     passwords,
		"PasswordCount": len(passwords),
	}

	if err := tmpl.ExecuteTemplate(w, "stats.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
