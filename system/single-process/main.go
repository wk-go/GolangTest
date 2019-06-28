package main

// 控制程序之启动单一进程
// 参考:https://golangtc.com/t/56342908b09ecc3ac5000052

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	DS = string(os.PathSeparator)
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	start(&wg)
	wg.Wait()
}

func start(wg *sync.WaitGroup) {
	defer wg.Done()
	iManPid := fmt.Sprint(os.Getpid())
	tmpDir := os.TempDir()

	if err := ProcExsit(tmpDir); err == nil {
		pidFile, _ := os.Create(tmpDir + DS + "single-process-test.pid")
		fmt.Println("pidFile:", pidFile.Name())
		defer pidFile.Close()

		pidFile.WriteString(iManPid)
	} else {
		log.Println(err)
		os.Exit(1)
	}

	//阻塞
	for {
		log.Println("Running")
		time.Sleep(10 * time.Second)
	}
}

// 判断进程是否启动
func ProcExsit(tmpDir string) (err error) {
	iManPidFile, err := os.Open(tmpDir + "\\single-process-test.pid")
	defer iManPidFile.Close()

	if err == nil {
		filePid, err := ioutil.ReadAll(iManPidFile)
		if err == nil {
			pidStr := fmt.Sprintf("%s", filePid)
			pid, _ := strconv.Atoi(pidStr)
			_, err := os.FindProcess(pid)
			if err == nil {
				return errors.New("[ERROR] 程序已启动.")
			}
		}
	}
	return nil
}
