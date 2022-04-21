package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := generateHash([]byte("Hello World!"))
	fmt.Println("hash:" + hash)

	hash2 := generateHash2([]byte("Hello World!"))
	fmt.Println("hash2:" + hash2)
}

func generateHash(message []byte) string {
	bytes := sha256.Sum256(message)
	code := hex.EncodeToString(bytes[:])
	return code
}

func generateHash2(message []byte) string {
	// 创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	// 计算哈希值
	bytes := hash.Sum(nil)
	// 将字符串编码为16进制格式,返回字符串
	code := hex.EncodeToString(bytes)
	// 返回哈希值
	return code
}
