package data

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Board 裡存放一個 Ptt 看板的資料
type Board struct {
	name     string
	articles []*article
}

// Article 裡存放一個 Ptt 文章的資料
type article struct {
	title  string
	author string
	date   string
	likes  int
}

// AddArticle 會新增一篇文章到看板裡
func (b *Board) AddArticle(t, a, d string, l int) {
	b.articles = append(b.articles, &article{t, a, d, l})
}

// ExportAsXlsx 會將整個看板資料 b 匯出為一個檔名為 fn 的 xlsx 檔案
func (b *Board) ExportAsXlsx(fn string) error {
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

// NumberOfArticles 會回傳現在已經有幾筆文章資料
func (b *Board) NumberOfArticles() int {
	return len(b.articles)
}
