package main
import(
	"github.com/therecipe/qt/widgets"
	"os"
)


func main(){
	widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("hello world");
	window.Show()
	widgets.QApplication_Exec()
}