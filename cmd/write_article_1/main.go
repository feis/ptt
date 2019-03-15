package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    const bn = "電影版"

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(bn)

	xlsx.SetCellValue(bn, "A1", "文章標題")
	xlsx.SetCellValue(bn, "B1", "作者")
	xlsx.SetCellValue(bn, "C1", "日期")
	xlsx.SetCellValue(bn, "D1", "讚數")

	xlsx.SetColWidth(bn, "A", "A", 20)
	xlsx.SetColWidth(bn, "B", "B", 10)

    xlsx.SetCellValue(bn, "A2", "假標題")
    xlsx.SetCellValue(bn, "B2", "假作者")
    xlsx.SetCellValue(bn, "C2", "3/14")
    xlsx.SetCellValue(bn, "D2", "5")

	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	xlsx.SaveAs("output.xlsx")
}
