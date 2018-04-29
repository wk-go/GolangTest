package main

import (
    "fmt"

    "github.com/360EntSecGroup-Skylar/excelize"
    "strconv"
)
func main(){
    Write()
    Read()
}

func Write() {
    xlsx := excelize.NewFile()
    // Create a new sheet.
    index := xlsx.NewSheet("Sheet2")
    // Set value of a cell.
    xlsx.SetCellValue("Sheet2", "A1", "Hello world.")
    // Set active sheet of the workbook.
    xlsx.SetActiveSheet(index)


    xlsx.SetCellValue("Sheet1", "A1", "Label1")
    xlsx.SetCellValue("Sheet1", "B1", "Label2")
    for i:=2; i<100; i++{
        k :=strconv.Itoa(i)
        xlsx.SetCellValue("Sheet1", "A"+k, "key_"+k)
        xlsx.SetCellValue("Sheet1", "B"+k, "val_"+k)
    }


    // Save xlsx file by the given path.
    err := xlsx.SaveAs("./Book1.xlsx")
    if err != nil {
        fmt.Println(err)
    }
}

func Read() {
    xlsx, err := excelize.OpenFile("./Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get value from cell by given worksheet name and axis.
    cell := xlsx.GetCellValue("Sheet1", "B2")
    fmt.Println(cell)
    // Get all the rows in the Sheet1.
    rows := xlsx.GetRows("Sheet1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}