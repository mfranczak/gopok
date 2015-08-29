package main

import (
	"net/http"
	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/new-game", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("name")))
	rw.Write(markdown)
}

