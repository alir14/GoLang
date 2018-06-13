package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

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

// func main() {
// 	var siteMap SitemapIndex
// 	var news News
// 	news_map := make(map[string]NewsMap)

// 	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	xml.Unmarshal(bytes, &siteMap)

// 	for _, Location := range siteMap.Locations {
// 		resp, _ := http.Get(Location)
// 		bytes, _ := ioutil.ReadAll(resp.Body)
// 		xml.Unmarshal(bytes, &news)

// 		for idx, _ := range news.Titles {
// 			news_map[news.Titles[idx]] = NewsMap{news.Keywords[idx], news.Locations[idx]}
// 		}
// 	}
// 	for idx, data := range news_map {
// 		fmt.Println("\n", idx)
// 		fmt.Println("\n", data.Keyword)
// 		fmt.Println("\n", data.Location)
// 	}

// }
