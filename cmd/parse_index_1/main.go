package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://www.ptt.cc/bbs/movie/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	// #main-container > div.r-list-container.action-bar-margin.bbs-screen > div:nth-child(2) > div.title > a
	doc.Find("#main-container > div.r-list-container.action-bar-margin.bbs-screen > div.r-ent > div.title > a").
		Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})
}
