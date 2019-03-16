package data1

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ExportAsXlsx 將看板資料 b 匯出成檔名為 fn 的 xlsx 檔案
func ExportAsXlsx(b *Board, fn string) error {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(b.Name)

	xlsx.SetCellValue(b.Name, "A1", "文章標題")
	xlsx.SetCellValue(b.Name, "B1", "作者")
	xlsx.SetCellValue(b.Name, "C1", "日期")
	xlsx.SetCellValue(b.Name, "D1", "讚數")

	xlsx.SetColWidth(b.Name, "A", "A", 60)
	xlsx.SetColWidth(b.Name, "B", "B", 20)

	for i, a := range b.Articles {
		xlsx.SetCellValue(b.Name, fmt.Sprintf("A%d", i+2), a.Title)
		xlsx.SetCellValue(b.Name, fmt.Sprintf("B%d", i+2), a.Author)
		xlsx.SetCellValue(b.Name, fmt.Sprintf("C%d", i+2), a.Date)
		xlsx.SetCellValue(b.Name, fmt.Sprintf("D%d", i+2), a.Likes)
	}

	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	return xlsx.SaveAs("output.xlsx")
}
