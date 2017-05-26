package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "time"
    "os"
)

func main() {
    var targetUrlLineEdit,sourceFileLineEdit, targetFileNameLineEdit *walk.LineEdit
    var targetUrlLabel,sourceFileLabel, targetFileNameLabel *walk.Label
    var selectFileButton, clearFileButton, downloadButton *walk.PushButton
    var logTextEdit *walk.TextEdit
    var labelMinSize = Size{Width:70}
    var mw = &walk.MainWindow{}
    MainWindow{
        AssignTo: &mw,
        Title: "Web Page Downloader",
        Layout: VBox{},
        MinSize: Size{600, 400},
        Children:[]Widget{
            Composite{
                Layout:HBox{},
                Children:[]Widget{
                    Label{AssignTo:&targetUrlLabel,Text:"目标网址:",MinSize:labelMinSize},
                    LineEdit{AssignTo:&targetUrlLineEdit},
                },
            },
            Composite{
                Layout:HBox{},
                Children: []Widget{
                    Label{AssignTo:&sourceFileLabel,Text:"本地源文件:",MinSize:labelMinSize},
                    LineEdit{AssignTo:&sourceFileLineEdit},
                    PushButton{AssignTo:&selectFileButton,Text:"浏览",OnClicked:func(){
                        selectFile(sourceFileLineEdit,mw)
                    }},
                    PushButton{AssignTo:&clearFileButton, Text:"清除", OnClicked:func(){
                        clearFile(sourceFileLineEdit)
                    }},
                },
            },
            Composite{
                Layout:HBox{},
                Children:[]Widget{
                    Label{AssignTo:&targetFileNameLabel,Text:"保存文件名:",MinSize:labelMinSize},
                    LineEdit{AssignTo:&targetFileNameLineEdit},
                },
            },
            Composite{
                Layout:HBox{},
                Children: []Widget{
                    PushButton{AssignTo:&downloadButton, Text:"下载", OnClicked:func(){
                        downloadButton.SetEnabled(false)
                        download(logTextEdit)
                        time.Sleep(2*time.Second)
                        downloadButton.SetEnabled(true)
                    }},
                },
            },
            Composite{
                Layout:HBox{},
                Children: []Widget{
                    TextEdit{AssignTo:&logTextEdit,ReadOnly:true,VScroll:true},
                },
            },
        },
    }.Run()
}

func selectFile(showEdit *walk.LineEdit,mw *walk.MainWindow){
    showEdit.SetText("hello world");
    currPath,_ := os.Getwd()
    f := walk.FileDialog{Title:"选择源文件", InitialDirPath: currPath}
    if o,_ := f.ShowOpen(mw); o{
        showEdit.SetText(f.FilePath)
    }
}
func clearFile(showEdit *walk.LineEdit){
    showEdit.SetText("")
}

func download(logEdit *walk.TextEdit){
    logEdit.AppendText("start download\n")
}
