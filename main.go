package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

type asciiArt struct {
	AsciiText string
}

func main() {
	fileserver := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileserver)
	http.HandleFunc("/ascii-art", InfoHandle)

	fmt.Printf("Starting server on port :9800 \n")
	fmt.Printf("Use 👉 Control+C to stop server \n")

	http.ListenAndServe(":9800", nil)
}

func InfoHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error Page 404 not found", 404)
		return
	}
	text := r.FormValue("textvalue")
	banner := r.FormValue("banner")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	bannerLines, _ := readBannerFile("Banners/" + banner + ".txt")
	asciiart := string(GenerateASCIIArt(text, bannerLines))
	structData := asciiArt{
		AsciiText: asciiart,
	}

	tpl, _ = template.ParseFiles("templates/index.html")
	err := tpl.Execute(w, structData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func readBannerFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var bannerLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bannerLines = append(bannerLines, scanner.Text())
	}
	return bannerLines, scanner.Err()
}
func GenerateASCIIArt(inputString string, bannerLines []string) string {
	var asciiArt strings.Builder
	words := strings.Split(inputString, `\n`)
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				asciiArt.WriteString("\n")
			}
			continue
		}
		for row := 1; row < 9; row++ {
			for _, char := range word {
				lineIndex := (int(char)-32)*9 + row
				if lineIndex >= 0 && lineIndex < len(bannerLines) {
					asciiArt.WriteString(bannerLines[lineIndex])
				}
			}
			asciiArt.WriteString("\n")
		}
		if i < len(words)-1 {
			asciiArt.WriteString("\n")
		}
	}
	return asciiArt.String()
}
