//+build ignore

package main

import (
	"fmt"
	"github.com/gonutz/wui/v2"
)

func main() {
	windowFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	window := wui.NewWindow()
	window.SetFont(windowFont)
	window.SetTitle("Window")

	textEdit1 := wui.NewTextEdit()
	textEdit1.SetBounds(43, 26, 150, 50)
	textEdit1.SetText("Text Edit")
	window.Add(textEdit1)

	editLine1 := wui.NewEditLine()
	editLine1.SetBounds(42, 110, 150, 20)
	editLine1.SetText("Text Edit Line")
	window.Add(editLine1)

	button1 := wui.NewButton()
	button1.SetBounds(70, 207, 85, 25)
	button1.SetText("Button")
	button1.SetOnClick(func() {
		fmt.Println("Hello world")
		dialog := wui.NewFileOpenDialog()
		dialog.SetTitle("选择一个文件")
		dialog.ExecuteSingleSelection(window)
	})
	window.Add(button1)

	window.Show()
}
