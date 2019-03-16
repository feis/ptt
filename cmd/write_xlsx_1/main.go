package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	xlsx := excelize.NewFile()
	xlsx.SaveAs("output.xlsx")
}
