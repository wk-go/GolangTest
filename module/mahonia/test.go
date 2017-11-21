package main

import (
    "github.com/axgle/mahonia"
    "fmt"
)



func main() {
    //"你好，世界！"的GBK编码
    testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
    var testStr string
    utfStr := "你好，世界！"
    testStr = string(testBytes)



    var dec mahonia.Decoder
    dec = mahonia.NewDecoder("gbk")
    fmt.Println("GBK to UTF-8: ", dec.ConvertString(testStr), " bytes:", testBytes)




    var enc mahonia.Encoder
    enc = mahonia.NewEncoder("gbk")
    ret := enc.ConvertString(utfStr)
    fmt.Println("UTF-8 to GBK: ", ret , " bytes: ", string(testBytes))

    return

}
