package main

import (
	"regexp"
	"fmt"
)

func main(){

	//domain test
	domainPattern := "([-a-z0-9A-Z].*?)\\.([-a-z0-9A-Z].*?)\\.(com|top|site|com\\.cn)"
	fmt.Println(regexp.MatchString(domainPattern,"baidu.com"))
	fmt.Println(regexp.MatchString(domainPattern,"www.baidu.com"))
	fmt.Println(regexp.MatchString(domainPattern,"tieba.baidu.com"))
	fmt.Println(regexp.MatchString(domainPattern,"www.sina.com.cn"))
	fmt.Println(regexp.MatchString(domainPattern,"blog.sina.com.cn"))
}