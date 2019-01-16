package main

import (
"fmt"
"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()


	c.OnHTML("div.manga-chapter-image", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		fmt.Println("---------")
	})


	c.Visit("https://comicvn.net/truyen-tranh-online/iron-wok-jan-dau-bep-sieu-dang/chapter-1-316596")
}