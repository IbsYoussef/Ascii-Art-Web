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

// enableCORS is a middleware to enable CORS for all routes
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
