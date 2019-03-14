package main

import data "github.com/feis/ptt/pkg/data1"

func main() {
	b := &data.Board{
		Name: "電影版",
		Articles: []data.Article{
			data.Article{
				Title:  "測試用的標題",
				Author: "我也不知道",
				Date:   "3/14",
				Likes:  -5,
			},
		},
	}

	data.ExportAsXlsx(b, "use_lib.xlsx")
}
