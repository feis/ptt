package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excelize.NewFile()
	err := xlsx.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
