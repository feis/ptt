package main

import data "github.com/feis/ptt/pkg/data2"

func main() {
	b := data.NewBoard("電影版")

	b.AddArticle("測試用的標題", "我也不知道", "3/14", -5)
	b.AddArticle("第二篇也試試看", "是我啦", "2/14", 10)

	b.ExportAsXlsx("use_lib_2.xlsx")
}
