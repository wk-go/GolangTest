# 解读student.proto

## 逐行解读student.proto

1. protobuf 有2个版本，默认版本是 proto2，如果需要 proto3，则需要在非空非注释第一行使用 syntax = "proto3" 标明版本。
2. package，即包名声明符是可选的，用来防止不同的消息类型有命名冲突。
3. 消息类型 使用 message 关键字定义，Student 是类型名，name, male, scores 是该类型的 3 个字段，类型分别为 string, bool 和 []int32。字段可以是标量类型，也可以是合成类型。
4. 每个字段的修饰符默认是 singular，一般省略不写，repeated 表示字段可重复，即用来表示 Go 语言中的数组类型。
5. 每个字符 =后面的数字称为标识符，每个字段都需要提供一个唯一的标识符。标识符用来在消息的二进制格式中识别各个字段，一旦使用就不能够再改变，标识符的取值范围为 [1, 2^29 - 1] 。
6. .proto 文件可以写注释，单行注释 //，多行注释 /* ... */
7. 一个 .proto 文件中可以写多个

## 编译
`protoc --go_out=. *.proto`

注意`option go_package = ".;data";`要指定golang的包