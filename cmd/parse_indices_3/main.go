package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	data "github.com/feis/ptt/pkg/data2"
)

func download(url string, b *data.Board) (err error) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return
	}

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

	return
}

func main() {
	b := data.NewBoard("電影板")

	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("https://www.ptt.cc/bbs/movie/index%d.html", i)
		err := download(url, b)

		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Number of articles:", b.NumberOfArticles())

	err := b.ExportAsXlsx("output.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}
}
