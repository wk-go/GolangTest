package main
//回显错误信息
import (
    "bytes"
    "fmt"
    "os/exec"
    "runtime"
    "github.com/henrylee2cn/mahonia"
)
func main() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows"{
        cmd = exec.Command("cmd","/c","fail")
    }else{
        cmd = exec.Command("/bin/sh", "-c", "fail")
    }
    var stdErr bytes.Buffer
    cmd.Stderr = &stdErr
    if err := cmd.Run(); err != nil {
        fmt.Println("Run returns:", err)
    }
    if runtime.GOOS == "windows"{
        fmt.Println("Stderr w:", Gbk2Utf8(stdErr.String()))
    }else{
        fmt.Println("Stderr o:", stdErr.String())
    }
}

func Gbk2Utf8(s string)(string){
    dec:=mahonia.NewDecoder("gbk")
    return dec.ConvertString(s)
}