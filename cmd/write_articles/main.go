package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type article struct {
	title  string
	author string
	date   string
	likes  int
}

func exportAsXlsx(bn string, as []*article, fn string) error {
	xlsx := excelize.NewFile()
	xlsx.NewSheet(bn)

	xlsx.SetCellValue(bn, "A1", "文章標題")
	xlsx.SetCellValue(bn, "B1", "作者")
	xlsx.SetCellValue(bn, "C1", "日期")
	xlsx.SetCellValue(bn, "D1", "讚數")

	xlsx.SetColWidth(bn, "A", "A", 60)
	xlsx.SetColWidth(bn, "B", "B", 20)

	for i, a := range as {
		xlsx.SetCellValue(bn, fmt.Sprintf("A%d", i+2), a.title)
		xlsx.SetCellValue(bn, fmt.Sprintf("B%d", i+2), a.author)
		xlsx.SetCellValue(bn, fmt.Sprintf("C%d", i+2), a.date)
		xlsx.SetCellValue(bn, fmt.Sprintf("D%d", i+2), a.likes)
	}

	xlsx.DeleteSheet("Sheet1")
	return xlsx.SaveAs("output.xlsx")
}

func main() {
	bn := "電影版"
	as := []*article{
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
	}
	exportAsXlsx(bn, as, "output.xlsx")
}
