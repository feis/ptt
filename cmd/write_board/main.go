package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type board struct {
	name     string
	articles []*article
}

type article struct {
	title  string
	author string
	date   string
	likes  int
}

func exportAsXlsx(b *board, fn string) error {
	xlsx := excelize.NewFile()
	xlsx.NewSheet(b.name)

	xlsx.SetCellValue(b.name, "A1", "文章標題")
	xlsx.SetCellValue(b.name, "B1", "作者")
	xlsx.SetCellValue(b.name, "C1", "日期")
	xlsx.SetCellValue(b.name, "D1", "讚數")

	xlsx.SetColWidth(b.name, "A", "A", 60)
	xlsx.SetColWidth(b.name, "B", "B", 20)

	for i, a := range b.articles {
		xlsx.SetCellValue(b.name, fmt.Sprintf("A%d", i+2), a.title)
		xlsx.SetCellValue(b.name, fmt.Sprintf("B%d", i+2), a.author)
		xlsx.SetCellValue(b.name, fmt.Sprintf("C%d", i+2), a.date)
		xlsx.SetCellValue(b.name, fmt.Sprintf("D%d", i+2), a.likes)
	}

	xlsx.DeleteSheet("Sheet1")
	return xlsx.SaveAs(fn)
}

func main() {
	b := &board{
		name: "電影版",
		articles: []*article{
			&article{
				title:  "假標題",
				author: "假作者",
				date:   "3/14",
				likes:  5,
			},
			&article{
				title:  "第二篇",
				author: "不知道",
				date:   "3/15",
				likes:  -10,
			},
		},
	}
	err := exportAsXlsx(b, "output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
