package main

import (
	"encoding/hex"
	"fmt"
	helper "golang_test/module/crypto/rsa/helper"
)

func main() {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey := helper.Generate()
	fmt.Println("-------------------------------获取RSA私钥-----------------------------------------")
	fmt.Println(string(prvKey))
	fmt.Println("-------------------------------获取RSA公钥-----------------------------------------")
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData := helper.Signature([]byte(data), prvKey)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if helper.Verify([]byte(data), signData, pubKey) {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := helper.Encrypt([]byte(data), pubKey)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := helper.Decrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
}
