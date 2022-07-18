package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

//暴力破解RAR文件
var password = make(chan string) //创建管道，接收密码
var isOver = make(chan struct{})
var WG sync.WaitGroup

func main() {
	var rarFile, passwordFile, outPath string
	flag.StringVar(&rarFile, "f", "", "rar file path")
	flag.StringVar(&passwordFile, "password-file", "", "password file path")
	flag.StringVar(&outPath, "o", "./test", "rar file path")
	flag.Parse()

	if passwordFile != "" {
		passwdFile := NewPasswordFile(passwordFile)
		go passwdFile.GetPassword(password)
	}
	if outPath != "" {
		dir, err := ioutil.ReadDir(outPath)
		if os.IsNotExist(err) {
			os.MkdirAll(outPath, 0755)
			dir, _ = ioutil.ReadDir(outPath)
		}
		for _, d := range dir {
			os.RemoveAll(path.Join(outPath, d.Name()))
		}
	}
	if rarFile == "" {
		return
	}

Loop:
	for {
		select {
		case pw := <-password:
			go Decipher(rarFile, pw, outPath)
		case <-isOver:
			break Loop
		}
	}
	WG.Wait()
}

func Decipher(rarFile, password, outPath string) {
	WG.Add(1)
	cmd := exec.Command("unrar", "e", "-p"+password, rarFile, outPath) //解压出来保存 D/test 上
	out, _ := cmd.Output()

	fmt.Println("Try:", password)
	if strings.Contains(string(out), "All OK") { //len 248 为成功，每个人不同
		fmt.Printf("The password is：  %s \n", password)
		go func() { isOver <- struct{}{} }()
	}
	WG.Done()
}

type PasswordFile struct {
	Filepath string
}

func NewPasswordFile(filepath string) *PasswordFile {
	return &PasswordFile{
		Filepath: filepath,
	}
}

func (p *PasswordFile) GetPassword(password chan string) {
	fp, _ := os.OpenFile(p.Filepath, os.O_RDONLY, 6)
	defer fp.Close()

	r := bufio.NewReader(fp)
	for {
		pass, _, err2 := r.ReadLine()
		if err2 == io.EOF {
			break
		}
		password <- string(pass)
	}
	isOver <- struct{}{}
}

type Password struct {
}
