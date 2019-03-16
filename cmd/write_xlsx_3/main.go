package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Main")
	xlsx.SetActiveSheet(index)
	xlsx.DeleteSheet("Sheet1")
	err := xlsx.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
