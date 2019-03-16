package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excelize.NewFile()
	xlsx.NewSheet("Main")
	xlsx.DeleteSheet("Sheet1")
	err := xlsx.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
