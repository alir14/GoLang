package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"html/template"
// 	"io/ioutil"
// 	"net/http"
// 	"sync"
// )

// var waitGroup sync.WaitGroup

// type SitemapIndex struct {
// 	Locations []string `xml:"sitemap>loc"`
// }

// type News struct {
// 	Titles    []string `xml:"url>news>title"`
// 	Keywords  []string `xml:"url>news>keywords"`
// 	Locations []string `xml:"url>loc"`
// }

// type NewsMap struct {
// 	Keyword  string
// 	Location string
// }

// type NewsAggPage struct {
// 	Title string
// 	News  map[string]NewsMap
// }

// func main() {
// 	http.HandleFunc("/", index_Handler)
// 	http.HandleFunc("/agg/", newsAggHandler)
// 	http.ListenAndServe(":8000", nil)
// }

// func index_Handler(rw http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(rw, "<h1>Index Page</h1>")
// }

// func newsRoutine(channel chan News, Location string) {
// 	defer waitGroup.Done()
// 	var news News
// 	resp, _ := http.Get(Location)
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	xml.Unmarshal(bytes, &news)
// 	resp.Body.Close()
// 	channel <- news
// }

// func (siteMap *SitemapIndex) getNewsFeed() {
// 	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	xml.Unmarshal(bytes, &siteMap)
// 	resp.Body.Close()
// }

// func newsAggHandler(rw http.ResponseWriter, r *http.Request) {
// 	var siteMap SitemapIndex
// 	news_map := make(map[string]NewsMap)
// 	siteMap.getNewsFeed()
// 	queue := make(chan News, 30)
// 	for _, Location := range siteMap.Locations {
// 		waitGroup.Add(1)
// 		go newsRoutine(queue, Location)
// 	}

// 	waitGroup.Wait()
// 	close(queue)

// 	for element := range queue {
// 		for idx, _ := range element.Titles {
// 			news_map[element.Titles[idx]] = NewsMap{element.Keywords[idx], element.Locations[idx]}
// 		}
// 	}

// 	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
// 	t, _ := template.ParseFiles("newsTemplate.html")
// 	fmt.Println(t.Execute(rw, p))
// }
