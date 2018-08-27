package main

import (
	"regexp"
	"fmt"
)

func main(){
	re := regexp.MustCompile("ab?")
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo") == nil)
}