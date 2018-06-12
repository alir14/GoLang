package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// type SitemapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// }

// type Location struct {
// 	Loc string `xml:"loc"`
// }

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

func main() {
	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	//resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	for _, Item := range s.Locations {
		resp, _ := http.Get(Item)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &s)
	}

}

// func getData() string {

// }
