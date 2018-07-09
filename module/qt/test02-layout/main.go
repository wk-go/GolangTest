package main
import(
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"os"
)


func main(){
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.Resize(core.NewQSize2(600,400))
	window.SetWindowTitle("hello world");
	window.Show()

	widget := widgets.NewQWidget(window,0)

	//button
	button := widgets.NewQPushButton(window)
	button.Move(core.NewQPoint2(100,100))
	button.SetText("Click Me!")
	button.ConnectClicked(func(_ bool) {
		msgBox := widgets.NewQMessageBox(nil)
		msgBox.SetWindowTitle("信息")
		msgBox.SetText("确定吗？")
		msgBox.SetIcon(widgets.QMessageBox__Warning)
		msgBox.Exec()
	})
	//button.Show()

	hBox := widgets.NewQHBoxLayout2(widget)
	hBox.AddWidget(button, 0, 0)

	//似乎是必须要通过一个widget将layout添加到MainWindow
	window.SetCentralWidget(widget)

	widgets.QApplication_Exec()
}