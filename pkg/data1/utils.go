package data

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ExportAsXlsx 會將整個看板資料 b 匯出為一個檔名為 fn 的 xlsx 檔案
func ExportAsXlsx(b *Board, fn string) {
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
	xlsx.SaveAs(fn)
}
