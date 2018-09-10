package main

import (
	"regexp"
	"fmt"
)

func main(){
	re := regexp.MustCompile("fo.?")
	fmt.Printf("%q\n", re.FindString("seafood"))
	fmt.Printf("%q\n", re.FindString("meat"))
}