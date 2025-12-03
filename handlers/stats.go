package handlers

import (
	"net/http"
	"password-checker/services"
)

// store is the shared Storage instance, set in main.go
var store *services.Storage

// Func lets main.go provide the shared Storage instance
func InitStore(s *services.Storage) {
	store = s
}

// Stats shows all stored passwords and basic counts
func Stats(w http.ResponseWriter, r *http.Request) {
	// Uses the Storage methods
	checked := store.Checked          // Gets the list of checked passwords directly from the store slice
	generated := store.GetGenerated() // Get the list of generated passwords using the store method

	// Prepare data for the template
	data := map[string]interface{}{
		"Checked":        checked,
		"Generated":      generated,
		"CheckedCount":   len(checked),
		"GeneratedCount": len(generated),
	}

	// Show stats.html
	if err := tmpl.ExecuteTemplate(w, "stats.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
