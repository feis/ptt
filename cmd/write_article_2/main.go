package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

type article struct {
	title  string
	author string
	date   string
	likes  int
}

func exportAsXlsx(bn string, a *article, fn string) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(bn)

	xlsx.SetCellValue(bn, "A1", "文章標題")
	xlsx.SetCellValue(bn, "B1", "作者")
	xlsx.SetCellValue(bn, "C1", "日期")
	xlsx.SetCellValue(bn, "D1", "讚數")

	xlsx.SetColWidth(bn, "A", "A", 20)
	xlsx.SetColWidth(bn, "B", "B", 10)

	xlsx.SetCellValue(bn, "A2", a.title)
	xlsx.SetCellValue(bn, "B2", a.author)
	xlsx.SetCellValue(bn, "C2", a.date)
	xlsx.SetCellValue(bn, "D2", a.likes)

	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	xlsx.SaveAs("output.xlsx")
}

func main() {
	const bn = "電影版"

	a := &article{
		title:  "假標題",
		author: "假作者",
		date:   "3/14",
		likes:  5,
	}

	exportAsXlsx(bn, a, "output.xlsx")
}
