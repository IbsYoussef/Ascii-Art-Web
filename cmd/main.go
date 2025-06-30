package main

import (
	"ascii-art-web/internal/web"
	"fmt"
	"net/http"
)

func main() {
	web.SetupRoutes()

	fmt.Println("🚀 Server running at http://localhost:9800")
	fmt.Println("📂 Visit: http://localhost:9800")
	http.ListenAndServe(":9800", nil)
}
