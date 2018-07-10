package main
import(
	"github.com/therecipe/qt/widgets"
	"os"
	"strings"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/core"
)

func commandLineUrlArgument() *core.QUrl{
	args := os.Args
	for _,arg := range args{
		index := strings.Index(arg, "-")
		if index != 0{
			return core.QUrl_FromUserInput(arg)
		}
	}
	return core.QUrl_FromUserInput("http://www.qt.io")
}

func main(){
	widgets.NewQApplication(len(os.Args), os.Args)

	webView := webengine.NewQWebEngineView(nil)
	webView.SetUrl(commandLineUrlArgument())
	webView.Resize(core.NewQSize2(1024, 750))
	webView.Show()

	widgets.QApplication_Exec()
}
