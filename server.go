package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("./templates"))

	style := http.FileServer(http.Dir("./templates"))
	http.Handle("/", style)
	http.HandleFunc("/asciiMaker", posthandler)
	// http.HandleFunc("/", posthandler)
	http.ListenAndServe(":8080", nil)
}

func posthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			http.Error(w, "ERROR-404\nPage not found(", http.StatusNotFound)
			return
		}
		templates.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" {
		text1 := r.FormValue("text")
		font := r.FormValue("fonts")
		text := ""

		if strings.Contains(text1, "\r\n") {
			text = strings.ReplaceAll(text1, "\r\n", "\\n")
		} else {
			text = text1
		}
		// check if user type in proper ascii art
		for _, v := range text {
			if !(v >= 32 && v <= 126) {
				http.Error(w, "ERROR-400\nBad request!", http.StatusBadRequest)
				return
			}
		}
		file, err := os.Open(font + ".txt")

		if err != nil {
			http.Error(w, "ERROR-400\nBad Request!! \nPlease make sure you select a font.", http.StatusBadRequest)
			return
		}

		defer file.Close()
		Scanner := bufio.NewScanner(file)

		var lines []string
		for Scanner.Scan() {
			lines = append(lines, Scanner.Text())
		}
		asciiCh := make(map[int][]string)
		dec := 31

		for _, line := range lines {
			if line == "" {
				dec++
			} else {
				asciiCh[dec] = append(asciiCh[dec], line)

			}
		}
		var c = ""
		for i := 0; i < len(text); i++ {
			if text[i] == 92 && text[i+1] == 110 {
				c = PrintArt(text[:i], asciiCh) + PrintArt(text[i+2:], asciiCh)
			}
		}
		if !strings.Contains(text, "\\n") {
			c = PrintArt(text, asciiCh)
		}

		templates.ExecuteTemplate(w, "index.html", c)
	}

}

func PrintArt(n string, y map[int][]string) string {
	a := []string{}
	for j := 0; j < len(y[32]); j++ {
		for _, letter := range n {
			a = append(a, y[int(letter)][j])
		}
		a = append(a, "\n")
	}
	b := strings.Join(a, "")
	return b
}
