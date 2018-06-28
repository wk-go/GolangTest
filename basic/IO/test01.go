package main
//命令行交互式 交互式程序

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)
var (
	str string
	num int
	in  string
)
func main(){
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入消息或指令>")
		in,_ = f.ReadString('\n') //按下回车则读入一行输入
		in = strings.Replace(in,"\r","",-1)
		fmt.Println("input string len:", len(in))
		if len(in) == 1{
			continue
		}
		fmt.Printf("您输入的是:%s", in)

		fmt.Sscan(in, &str, &num) //将输入格式化
		fmt.Printf("您输入的第一个参数是：%v,第二个参数是:%v.\n", str, num)

		if str == "stop" || str == "exit" || str == "quit"{
			fmt.Println("Bye ^_^")
			break
		}
	}
}
