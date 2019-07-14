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

func exportAsXlsx(bn string, a *article, fn string) error {
	xlsx := excelize.NewFile()
	xlsx.NewSheet(bn)

	xlsx.SetCellValue(bn, "A1", "文章標題")
	xlsx.SetCellValue(bn, "B1", "作者")
	xlsx.SetCellValue(bn, "C1", "日期")
	xlsx.SetCellValue(bn, "D1", "讚數")

	xlsx.SetColWidth(bn, "A", "A", 60)
	xlsx.SetColWidth(bn, "B", "B", 20)

	xlsx.SetCellValue(bn, "A2", a.title)
	xlsx.SetCellValue(bn, "B2", a.author)
	xlsx.SetCellValue(bn, "C2", a.date)
	xlsx.SetCellValue(bn, "D2", a.likes)

	xlsx.DeleteSheet("Sheet1")
	return xlsx.SaveAs(fn)
}

func main() {
	bn := "電影版"

	a := &article{
		title:  "假標題",
		author: "假作者",
		date:   "3/14",
		likes:  5,
	}

	err := exportAsXlsx(bn, a, "output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
