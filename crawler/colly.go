package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	c.OnHTML("div.tit.hot > a", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		fmt.Println(e.Attr("title"))
		fmt.Printf("%s%s", "https://comicvn.net", e.Attr("href"))
		fmt.Println(e.Attr("---------"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://comicvn.net/the-loai")
}