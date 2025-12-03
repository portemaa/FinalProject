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

/* Please change anything that sucks, this is just something we can use to visualize and get it done soon.

password-checker/
│
├── main.go              // Entry point: starts server, wires routes to handlers
│
├── handlers/            // HTTP endpoints (controllers)
│   ├── home.go          // Person 1: landing page, intro, educational info
│   ├── check.go         // Person 2: password input form, calls strength + improvement services
│   ├── generate.go      // Person 3: endpoint to generate random strong password
│   ├── about.go         // Person 3: endpoint to show project info + educational content
│   └── stats.go         // Person 1: endpoint to show user statistics (counts of Weak/Moderate/Strong)
│
├── services/            // Business logic (service layer)
│   ├── stats.go         // Person 1: track submissions, calculate distributions
│   ├── education.go     // Person 1: provide educational content about password hygiene
│   ├── strength.go      // Person 2: analyze password strength (length, variety, predictability)
│   ├── improve.go       // Person 2: suggest improvements for Weak/Moderate passwords
│   ├── generator.go     // Person 3: logic for generating random strong passwords
│   └── storage.go       // Person 3: save/retrieve passwords locally (map or file)
│
├── utils/               // Shared helpers/utilities
│   ├── security.go      // Person 2:Common security helpers (hashing, validation)
│   ├── evaluate.go      // Person 2:holds CheckPasswordStrength + CheckCharType helper functions
│
│
└── assets/              // Merged folder: HTML templates + static files
    ├── home.html        // Person 1: template for landing page
    ├── check.html       // Person 2: template for password input + feedback
    ├── generate.html    // Person 3: template for password generator
    ├── about.html       // Person 3: template for project info
    ├── style.css        // Shared: styling for all pages
    ├── script.js        // Shared: client-side interactivity
    └── 
*/
func main() {
	//Had to look up, REf: go.dev to help parse all templates in asset packages
	tmpl := template.Must(template.ParseGlob("assets/*.html"))
	// Create a shared memory storage instance
	store := services.NewMemStorage()

	// Initialize handlers with template and storage
	handlers.Init(tmpl)
	handlers.InitStore(store)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/check", handlers.Check)
	http.HandleFunc("/generate", handlers.Generate(tmpl, store))
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/stats", handlers.Stats)

	// This tells Go's web server to look inside the "assets" folder and deliver files directly to the browser.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", nil)

}
