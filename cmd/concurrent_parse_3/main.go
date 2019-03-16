package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	data "github.com/feis/ptt/pkg/data2"
)

type article struct {
	title  string
	author string
	date   string
	likes  int
}

func download(url string, c chan *article) {
	resp, err := http.Get(url)

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

			c <- &article{
				title:  t.Text(),
				author: a.Text(),
				date:   d.Text(),
				likes:  ls}
		})
}

func main() {
	start := time.Now()

	c := make(chan *article)
	q := make(chan bool)

	b := data.NewBoard("電影板")

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://www.ptt.cc/bbs/movie/index%d.html", i)
		go func() {
			download(url, c)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		q <- true
	}()

	for ok := true; ok; {
		select {
		case a := <-c:
			b.AddArticle(a.title, a.author, a.date, a.likes)
		case <-q:
			ok = false
		}
	}

	fmt.Println("Number of articles:", b.NumberOfArticles())

	// 因為存檔相當耗時且不穩定，所以我們就不計算存檔的時間
	fmt.Println("Elapsed time:", time.Since(start))

	err := b.ExportAsXlsx("output.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}
}
