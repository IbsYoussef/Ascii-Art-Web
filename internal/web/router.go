package web

import (
	"net/http"
)

// SetupRoutes configures all HTTP routes
func SetupRoutes() {
	// Serve static assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Serve index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	// Handle form submission
	http.HandleFunc("/ascii-art", AsciiHandler)
}
