package data

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Article 是存放文章資訊的結構
type Article struct {
	title  string
	author string
	date   string
	likes  int
}

// Board 是存放看板資訊的結構
type Board struct {
	name     string
	articles []Article
}

// AddArticle 會新增一篇文章到看板裡
func (b *Board) AddArticle(t, a, d string, l int) {
	b.articles = append(b.articles, Article{t, a, d, l})
}

// ExportAsXlsx 會將整個看板資料 b 匯出為一個檔名為 fn 的 xlsx 檔案
func (b *Board) ExportAsXlsx(fn string) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(b.name)

	xlsx.SetCellValue(b.name, "A1", "文章標題")
	xlsx.SetCellValue(b.name, "B1", "作者")
	xlsx.SetCellValue(b.name, "C1", "日期")
	xlsx.SetCellValue(b.name, "D1", "讚數")

	xlsx.SetColWidth(b.name, "A", "A", 20)
	xlsx.SetColWidth(b.name, "B", "B", 10)

	for i, a := range b.articles {
		xlsx.SetCellValue(b.name, fmt.Sprintf("A%d", i+2), a.title)
		xlsx.SetCellValue(b.name, fmt.Sprintf("B%d", i+2), a.author)
		xlsx.SetCellValue(b.name, fmt.Sprintf("C%d", i+2), a.date)
		xlsx.SetCellValue(b.name, fmt.Sprintf("D%d", i+2), a.likes)
	}

	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	xlsx.SaveAs(fn)
}
