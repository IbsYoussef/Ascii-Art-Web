package web

import (
	"ascii-art-web/internal/ascii"
	"html/template"
	"net/http"
)

type AsciiArt struct {
	AsciiText string
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != http.MethodPost {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("textvalue")
	banner := r.FormValue("banner")

	bannerMap, err := ascii.LoadBannerFile(banner)
	if err != nil {
		http.Error(w, "Failed to load banner", http.StatusInternalServerError)
		return
	}

	asciiResult := ascii.RenderAscii(text, bannerMap)

	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, AsciiArt{AsciiText: asciiResult})
}
