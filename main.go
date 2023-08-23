package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Formdata struct {
	data string
}

func main() {

	fileserver := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fileserver)
	http.HandleFunc("/process", handleInput)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func handleInput(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	value := r.FormValue("fname")
	banner := r.FormValue("styles")
	valueStruct := Formdata{
		data: value,
	}

	tpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, valueStruct)
	fmt.Fprintf(w, "%s ,%s", valueStruct.data, banner)

}
