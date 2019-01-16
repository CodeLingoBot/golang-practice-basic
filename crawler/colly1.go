package main

import (
    "fmt"
    "github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector()


    c.OnHTML("li.active > ul> li> a", func(e *colly.HTMLElement) {
        fmt.Println(e.Text)
        fmt.Printf("%s%s", "https://comicvn.net", e.Attr("href"))
        fmt.Println(e.Attr("---------"))
    })


    c.Visit("https://comicvn.net/truyen-tranh-online/iron-wok-jan-dau-bep-sieu-dang-13283")
}