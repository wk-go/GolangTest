package main

import (
	"log"
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

func main() {
	command := "ping"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	path, err := exec.LookPath(command)
	if err != nil {
		log.Printf("installing [%s] is in your future\n", command)
	}
	log.Printf("[%s] is available at %s\n", command, path)

	path1,err := GetCurrentPath(command)
	if err != nil {
		log.Printf("GetCurrentPath error: %s\n", err)
	}
	log.Printf("GetCurrentPath:%s\n", path1)
}

func GetCurrentPath(command string) (string, error) {
	file, err := exec.LookPath(command)
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`Error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}