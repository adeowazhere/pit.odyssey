package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func presentPage(fileName string, w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = page.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	default:
		presentPage("index.html", w, r)
	}
}

func main() {
	//http.Handle("/public/css/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}
