package main

import (
	"fmt"
	"html/template"
	"net/http"
	"password-checker/handlers"
	"password-checker/services"
)

//Revised Visual of Structure (I hope its fine I started the main, I also provided a compressed updated structure. Feel free to
// change anything that sucks or is annoying. The IDE I makes git chages in semi real time, so no need to push)

/*
	Please change anything that sucks, this is just something we can use to visualize and get it done soon.

password-checker/
│
├── main.go              // Entry point: starts server, wires routes to handlers
│
├── handlers/            // HTTP endpoints (controllers)
│   ├── home.go          // Landing page
│   ├── check.go         // Password input form, calls strength + improvement services
│   ├── generate.go      // Endpoint to generate random strong password
│   ├── about.go         // Endpoint to show project info
│   ├── stats.go         // Endpoint to show user statistics
│   └── init.go          // Initializes shared template and storage
│
├── services/            // Business logic (service layer)
│   ├── generator.go     // Logic for generating random strong passwords
│   ├── improve.go       // Suggest improvements for Weak/Moderate passwords
│   ├── strength.go      // Analyze password strength
│   └── storage.go       // Save/retrieve passwords locally (map with mutex)
│
├── utils/               // Shared helpers/utilities
│   ├── evaluator.go     // Holds CheckPasswordStrength + CheckCharType helper functions
│   └── security.go      // Common security helpers (hashing, validation)
│
└── assets/              // HTML templates + static files
    ├── home.html        // Template for landing page
    ├── check.html       // Template for password input + feedback
    ├── generate.html    // Template for password generator
    ├── about.html       // Template for project info
    ├── style.css        // Styling for all pages
    ├── script.js        // Client-side interactivity
    └── navbar.html      // Shared navigation bar across pages

*/
│
*/
func main() {
	//Had to look up, REf: go.dev to help parse all templates in asset packages
	tmpl := template.Must(template.ParseGlob("assets/*.html"))
	template.Must(tmpl.ParseGlob("assets/partials/*.html"))
	// Create a shared memory storage instance
	store := services.NewMemStorage()

	// Initialize handlers with template and storage
	handlers.Init(tmpl)
	handlers.InitStore(store)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/check", handlers.Check)
	http.HandleFunc("/generate", handlers.Generate(tmpl, store))
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/stats", handlers.Stats)

	// This tells Go's web server to look inside the "assets" folder and deliver files directly to the browser.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", nil)

}
