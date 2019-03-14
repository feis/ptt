package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type article struct {
	Title  string
	Author string
	Date   string
	Likes  int
}

type board struct {
	Name     string
	Articles []article
}

func exportAsXlsx(b *board, fn string) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(b.Name)

	xlsx.SetCellValue(b.Name, "A1", "文章標題")
	xlsx.SetCellValue(b.Name, "B1", "作者")
	xlsx.SetCellValue(b.Name, "C1", "日期")
	xlsx.SetCellValue(b.Name, "D1", "讚數")

	xlsx.SetColWidth(b.Name, "A", "A", 20)
	xlsx.SetColWidth(b.Name, "B", "B", 10)

	for i, a := range b.Articles {
		xlsx.SetCellValue(b.Name, fmt.Sprintf("A%d", i+2), a.Title)
		xlsx.SetCellValue(b.Name, fmt.Sprintf("B%d", i+2), a.Author)
		xlsx.SetCellValue(b.Name, fmt.Sprintf("C%d", i+2), a.Date)
		xlsx.SetCellValue(b.Name, fmt.Sprintf("D%d", i+2), a.Likes)
	}

	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	xlsx.SaveAs("fake_article.xlsx")
}

func main() {
	b := &board{
		Name: "電影版",
		Articles: []article{
			article{
				Title:  "測試用的標題",
				Author: "我也不知道",
				Date:   "3/14",
				Likes:  -5,
			},
		},
	}

	exportAsXlsx(b, "fake_article.xlsx")
}
