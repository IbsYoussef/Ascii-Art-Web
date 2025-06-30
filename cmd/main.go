package main

import (
	"ascii-art-web/internal/web"
	"fmt"
	"net/http"
	"os"
)

func main() {
	web.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
