package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	bn := "電影版"

	xlsx := excelize.NewFile()
	xlsx.NewSheet(bn)

	xlsx.SetCellValue(bn, "A1", "文章標題")
	xlsx.SetCellValue(bn, "B1", "作者")
	xlsx.SetCellValue(bn, "C1", "日期")
	xlsx.SetCellValue(bn, "D1", "讚數")

	xlsx.SetColWidth(bn, "A", "A", 60)
	xlsx.SetColWidth(bn, "B", "B", 20)

	xlsx.SetCellValue(bn, "A2", "假標題")
	xlsx.SetCellValue(bn, "B2", "假作者")
	xlsx.SetCellValue(bn, "C2", "3/14")
	xlsx.SetCellValue(bn, "D2", "5")

	xlsx.DeleteSheet("Sheet1")
	err := xlsx.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
