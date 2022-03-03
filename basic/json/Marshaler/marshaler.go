package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarshelerTest struct {
	json.Marshaler
}

func (m *MarshelerTest) MarshalJSON() ([]byte, error) {
	return []byte("{\"name\": \"world\"}"), nil
}

type Time2 time.Time

func (t Time2) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(t).Format("2006-01-02 15:04:05") + "\""), nil
}

type TimeInside struct {
	Name     string
	Datetime Time2
}

func main() {
	// 直接使用
	x := &MarshelerTest{}
	y, _ := json.Marshal(x)
	fmt.Printf("X:%s\n", y)

	//成员的修改
	x1 := TimeInside{Name: "x1", Datetime: Time2(time.Now())}
	y1, _ := json.Marshal(x1)
	fmt.Printf("X1:%s\n", y1)
}
