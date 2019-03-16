package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excelize.NewFile()
	xlsx.NewSheet("Main")
	xlsx.SetCellValue("Main", "A1", "Hello")
	xlsx.SetCellValue("Main", "B1", "world!")
	xlsx.DeleteSheet("Sheet1")
	err := xlsx.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
