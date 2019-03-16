package main

import (
	"fmt"

	data "github.com/feis/ptt/pkg/data1"
)

func main() {
	b := &data.Board{
		Name: "電影版",
		Articles: []*data.Article{
			&data.Article{
				Title:  "假標題",
				Author: "假作者",
				Date:   "3/14",
				Likes:  5,
			},
			&data.Article{
				Title:  "第二篇",
				Author: "不知道",
				Date:   "3/15",
				Likes:  -10,
			},
		},
	}
	err := data.ExportAsXlsx(b, "output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
