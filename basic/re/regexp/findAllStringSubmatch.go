package main

import (
	"regexp"
	"fmt"
)

func main(){
	re := regexp.MustCompile("a(x*)b")
	fmt.Printf("%q\n",re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
}