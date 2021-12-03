package main

import "fmt"

type Foo struct {
	ID   int64
	Name string
	//Bytes []byte            //不能比较
	//Map   map[string]string //不能比较
	//Func  func() error      //不能比较
}

func main() {
	var _testNil, _testNil2 map[string]string
	struct1 := Foo{ID: 1, Name: "Sam"}
	_struct1 := Foo{ID: 1, Name: "Sam"}
	struct2 := Foo{ID: 1, Name: "Peter"}

	var data = []map[string]interface{}{
		{"params": []interface{}{1, 1}, "expected": true},
		{"params": []interface{}{2, 2, 2, 2, 2}, "expected": true},
		{"params": []interface{}{1, 2}, "expected": false},
		{"params": []interface{}{"1", 1}, "expected": false},
		{"params": []interface{}{nil, nil}, "expected": true},
		{"params": []interface{}{&_testNil, &_testNil2}, "expected": false},
		{"params": []interface{}{nil, _testNil}, "expected": false},
		//{"params": []interface{}{_testNil, _testNil2}, "expected": false},//slices、maps、functions类型不能参与比较，同样包含这三种类型的结构体也不能参与比较
		{"params": []interface{}{struct1, _struct1}, "expected": true},
		{"params": []interface{}{struct1, struct2}, "expected": false},
		{"params": []interface{}{&struct1, &struct1}, "expected": true},
		{"params": []interface{}{&struct1, &_struct1}, "expected": false},
	}

	for _, v := range data {
		result := compare(v["params"].([]interface{})...)
		fmt.Printf("expected:%t,result:%t,params:%#v\n", v["expected"].(bool), result, v["params"].([]interface{}))
	}
}

func compare(params ...interface{}) bool {
	if len(params) < 2 {
		return false
	}

	for i, v := range params {
		if i == 0 {
			continue
		}
		if params[0] != v {
			return false
		}
	}
	return true
}
