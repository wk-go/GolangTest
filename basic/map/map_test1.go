/*
 map 测试脚本1
 */
package main

import(
	"fmt"
)
//article
type article struct{
	id	int32
	title string
	content string
}
func main(){
	//make 方式创建
	mapStrInt := make(map[string]int)
	mapStrInt["a"] = 1;
	mapStrInt["b"] = 2;
	mapStrInt["c"] = 3;
	mapStrInt["d"] = 4;
	mapStrInt["e"] = 5;
	for key, value := range mapStrInt{
		fmt.Printf("%v:%v\n",key,value);
	}
	fmt.Println()

	//复合结构定义方式
	mapArt := map[int]article{
		1:{1,"鬼吹灯","鬼吹灯简介"},
		2:{2,"盗墓笔记","盗墓笔记简介"},
		3:{3,"资治通鉴","资治通鉴简介"},
		4:{title:"周易",content:"周易简介"},
		5:{id:5,title:"西游记"},
	}
	for _, v := range mapArt{
		fmt.Printf("id:%v,name:%v,content:%v\n",v.id,v.title,v.content)
	}
	fmt.Println()

	//存在键名ok为true
	v,ok := mapArt[1]
	fmt.Printf("map nil: %v-%v\n",ok,v)
	//不存在键名ok为false
	v,ok = mapArt[100]
	fmt.Printf("map nil: %v-%v\n",ok,v)
}