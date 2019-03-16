package main

import (
	"fmt"

	data "github.com/feis/ptt/pkg/data2"
)

func main() {
	b := data.NewBoard("電影版")
	b.AddArticle("假標題", "假作者", "3/14", 5)
	b.AddArticle("第二篇", "不知道", "3/15", -10)

	err := b.ExportAsXlsx("output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
