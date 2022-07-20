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
	var rarFile, passwordFile, outPath, mode, symbols string
	var length, minLength, maxLength int
	flag.StringVar(&mode, "mode", "dict", "mode:dict,exhaustivity")
	flag.StringVar(&rarFile, "f", "", "rar file path")
	flag.StringVar(&passwordFile, "password-file", "", "password file path")
	flag.StringVar(&outPath, "o", "./test", "rar file path")
	flag.StringVar(&symbols, "symbols", "", "穷举字符")
	flag.IntVar(&length, "length", 5, "穷举模式下固定长度")
	flag.IntVar(&minLength, "min-length", 0, "穷举模式下长度区间最短长度")
	flag.IntVar(&maxLength, "max-length", 0, "穷举模式下长度区间最长长度")
	flag.Parse()

	if rarFile == "" {
		return
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

	// 密码字典模式
	if mode == "dict" {
		if passwordFile != "" {
			passwdFile := NewPasswordFile(passwordFile)
			go passwdFile.GetPassword(password)
		}
	}

	//穷举模式
	if mode == "exhaustivity" {
		pw := NewPassword()
		if symbols != "" {
			pw.SetSymbols(symbols)
		}
		var gen func(chan string)
		if length > 0 && minLength == 0 && maxLength == 0 {
			gen = pw.GeneratePassword(length)
		}
		if minLength > 0 && maxLength > 0 {
			pw.MinLength = minLength
			pw.MaxLength = maxLength
			gen = pw.GeneratePasswordRange()
		}
		go func() {
			gen(password)
			isOver <- struct{}{}
		}()
	}

	handle(rarFile, outPath)

	WG.Wait()
}

func handle(rarFile, outPath string) {
Loop:
	for {
		select {
		case pw := <-password:
			go Decipher(rarFile, pw, outPath)
		case <-isOver:
			break Loop
		}
	}
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
	Symbol    string
	symbolLen int
	Length    int
	MaxLength int
	MinLength int
}

func NewPassword() *Password {
	p := &Password{}
	p.SetSymbols("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
	return p
}

func (p *Password) GeneratePassword(length int) func(password chan string) {
	if length == 0 {
		length = p.Length
	}
	if length == 0 {
		return nil
	}
	return func(password chan string) {
		lenSlice := make([]int, length)
		symbolLen := p.symbolLen
		passwd := make([]byte, length)
		posFlag := 0
		breakFlag := 0
	For:
		for {
			breakFlag = 0
			for i, v := range lenSlice {

				passwd[i] = p.Symbol[v]
				if (v + 1) == symbolLen {
					breakFlag++
				}
				if breakFlag == len(lenSlice) {
					break For
				}
			}
			password <- string(passwd)
			for i, _ := range lenSlice {
				if i == 0 {
					lenSlice[i]++
					if lenSlice[i] >= symbolLen {
						posFlag = 1
						lenSlice[i] = lenSlice[i] % symbolLen
					}
					continue
				}
				if posFlag > 0 {
					lenSlice[i] += posFlag
					posFlag = 0
					if lenSlice[i] >= symbolLen {
						posFlag = 1
						lenSlice[i] = lenSlice[i] % symbolLen
					}
				}
			}
		}
	}
}

func (p *Password) GeneratePasswordRange() func(password chan string) {
	return func(password chan string) {
		for i := p.MinLength; i <= p.MaxLength; i++ {
			gen := p.GeneratePassword(i)
			gen(password)
		}
	}
}

func (p *Password) SetSymbols(s string) {
	if len(s) == 0 {
		return
	}
	p.Symbol = s
	p.symbolLen = len(s)
}
