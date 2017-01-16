package main

import (
    "fmt"
    "strings"
)

func main() {
    // 连接字符串
    s := []string{"foo", "bar", "baz"}
    fmt.Println("Join string:", strings.Join(s, "|"))

    // func Index(s, sep string) int 在字符串s中查找sep位置,返回位置值， 找不到返回-1
    fmt.Println("strings.Index(\"chicken\", \"ken\"):",strings.Index("chicken", "ken"))
    fmt.Println("strings.Index(\"chicken\", \"dmr\"):",strings.Index("chicken","dmr"))
    /* func Repeat(s string, count int) string
    重复s字符串count次，最后返回重复的字符串
    */
    fmt.Println("strings.Repeat(\"ab!\",10):", strings.Repeat("ab!",10))
    /*
    func Replace(s, old, new string, n int) string
    在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
     */

    fmt.Println("strings.Replace(\"Hello world!\",\"world\",\"china\",-1):", strings.Replace("Hello world!","world","china",-1))

    /*
    func Split(s, sep string) []string
    把s字符串按照sep分割，返回slice
     */
    fmt.Printf("%q\n", strings.Split("a,b,c", ","))
    fmt.Printf("%q\n",strings.Split("a man a plan a canal panama","a "))
    fmt.Printf("%q\n", strings.Split(" xyz ", ""))
    fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
    /*
    func Trim(s string, cutset string) string
    在s字符串的头部和尾部去除cutset指定的字符串
     */
    fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ","! "))
    /**
    func Fields(s string) []string
    去除s字符串的空格符，并且按照空格分割返回slice
     */
    fmt.Printf("Fields are: %q\n", strings.Fields("   Foo bar baz   "))
}
