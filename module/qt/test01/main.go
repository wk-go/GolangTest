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

	//button
	button := widgets.NewQPushButton(window)
	button.Move(core.NewQPoint2(100,100))
	button.SetText("Click Me!")
	button.Show()

	widgets.QApplication_Exec()
}