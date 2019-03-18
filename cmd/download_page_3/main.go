package main

import (
	"fmt"
	"io/ioutil"
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
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		if url, ok := s.Attr("href"); ok {
			s.SetAttr("href", "https:"+url)
		}
	})

	html, err := goquery.OuterHtml(doc.Selection)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("index.html", []byte(html), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
