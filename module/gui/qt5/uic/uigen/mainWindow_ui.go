// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIMainWindowMainWindow struct {
	Centralwidget        *widgets.QWidget
	VerticalLayoutWidget *widgets.QWidget
	VerticalLayout       *widgets.QVBoxLayout
	VerticalScrollBar    *widgets.QScrollBar
	PushButton           *widgets.QPushButton
	LineEdit             *widgets.QLineEdit
	Menubar              *widgets.QMenuBar
	Statusbar            *widgets.QStatusBar
}

func (this *UIMainWindowMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 800, 600))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.VerticalLayoutWidget = widgets.NewQWidget(this.Centralwidget, core.Qt__Widget)
	this.VerticalLayoutWidget.SetObjectName("VerticalLayoutWidget")
	this.VerticalLayoutWidget.SetGeometry(core.NewQRect4(180, 0, 611, 551))
	this.VerticalLayout = widgets.NewQVBoxLayout2(this.VerticalLayoutWidget)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)
	this.VerticalScrollBar = widgets.NewQScrollBar(this.VerticalLayoutWidget)
	this.VerticalScrollBar.SetObjectName("VerticalScrollBar")
	this.VerticalScrollBar.SetOrientation(core.Qt__Vertical)
	this.VerticalLayout.AddWidget(this.VerticalScrollBar, 0, 0)
	this.PushButton = widgets.NewQPushButton(this.Centralwidget)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetGeometry(core.NewQRect4(30, 20, 75, 23))
	this.LineEdit = widgets.NewQLineEdit(this.Centralwidget)
	this.LineEdit.SetObjectName("LineEdit")
	this.LineEdit.SetGeometry(core.NewQRect4(20, 60, 113, 20))
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 800, 23))
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)

	this.RetranslateUi(MainWindow)

}

func (this *UIMainWindowMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.PushButton.SetText(_translate("MainWindow", "开始", "", -1))
	this.LineEdit.SetText(_translate("MainWindow", "hello world!", "", -1))
}
