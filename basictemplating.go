package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

// func main() {
// 	http.HandleFunc("/", index_Handler)
// 	http.HandleFunc("/agg/", newsAggHandler)
// 	http.ListenAndServe(":8000", nil)
// }

func index_Handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "<h1>Index Page</h1>")
}

func newsAggHandler(rw http.ResponseWriter, r *http.Request) {

	var siteMap SitemapIndex
	var news News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &siteMap)

	for _, Location := range siteMap.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &news)

		for idx, _ := range news.Titles {
			news_map[news.Titles[idx]] = NewsMap{news.Keywords[idx], news.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("newsTemplate.html")
	fmt.Println(t.Execute(rw, p))
}
