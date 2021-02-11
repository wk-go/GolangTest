package main

//中文实体转换
import (
	"fmt"
	"html"
)

func main() {
	fmt.Println(html.EscapeString("中"))
	fmt.Printf("&#%d;\n", '中')
	s := fmt.Sprintf("&#%d;", '中')
	fmt.Printf("s:%s\n", s)
	s2 := html.UnescapeString(s)
	fmt.Printf("s2:%s\n", s2)
}
