package main

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"os"
	//"strings"
)

/**
压缩文件rar/test.rar的密码为4位数字密码。暴力破解。
实际密码为2563
*/
func main() {
	rarfile := archiver.NewRar()
	filename := "rar/test.rar"
	targetDir := "rar"
	for i := 0; i < 10000; i++ {
		os.Remove("rar/flag.txt")
		rarfile.Password = fmt.Sprintf("%04d", i)
		err := rarfile.Unarchive(filename, targetDir)
		if err == nil {
			fmt.Println("success:", rarfile.Password)
			break
		}
		fmt.Println("fail:", rarfile.Password, "error:", err)
		/*if strings.Index(err.Error(),"file already exists") >= 0{
		    os.Remove("rar/flag.txt")
		}*/
	}
}
