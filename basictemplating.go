package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// _ some varibale you do not use

type NewsAggPage struct {
	Title string
	News  string
}

func main() {
	http.HandleFunc("/", index_Handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

func index_Handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "<h1>Index Page</h1>")
}

func newsAggHandler(rw http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(rw, p)
}
