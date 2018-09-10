package main

import (
	"strings"
	"fmt"
	"regexp"
)

func main(){
	testDomains := []string{"m.test.tst", "x.m.test2.tst2", "x.testdomain.com.cn","a.b.c.testdomain.net.cn"}

	for _,domain := range testDomains{
		split1(domain)
	}
}

func split1(domain string){
	index := 2
	if matched,_ := regexp.MatchString("(com.cn|net.cn|org.cn|gov.cn)", domain); matched{
		index = 3
	}

	domainArray := strings.Split(domain, ".")
	fmt.Println(domainArray)
	fmt.Println(domainArray[len(domainArray)-index:])
	rootDomain := strings.Join(domainArray[len(domainArray)-index:],".")
	fmt.Println(rootDomain)
}

func split2(domain string){
	//lastIndex := strings.LastIndex(domain,".")
}