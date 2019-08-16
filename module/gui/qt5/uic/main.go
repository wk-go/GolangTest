package main

import (
	"github.com/therecipe/qt/widgets"
	"golang_test/module/gui/qt5/uic/uigen"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	ui := &uigen.UIMainWindowMainWindow{}
	ui.SetupUI(window)
	window.Show()
	app.Exec()
}
