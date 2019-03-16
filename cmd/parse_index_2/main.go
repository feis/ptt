package main

import (
	"fmt"
	"net/http"
	"strconv"

	data "github.com/feis/ptt/pkg/data2"

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

	b := data.NewBoard("電影板")

	doc.Find("#main-container > div.r-list-container.action-bar-margin.bbs-screen > div.r-ent").
		Each(func(i int, s *goquery.Selection) {
			t := s.Find("div.title > a").First()
			a := s.Find("div.meta > div.author").First()
			d := s.Find("div.meta > div.date").First()
			l := s.Find("div.nrec > span").First()

			ls, err := strconv.Atoi(l.Text())

			if err != nil {
				ls = 0
			}

			b.AddArticle(t.Text(), a.Text(), d.Text(), ls)
		})

	err = b.ExportAsXlsx("output.xlsx")

	if err != nil {
		fmt.Println(err)
	}
}
