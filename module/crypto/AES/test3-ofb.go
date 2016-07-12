package main
//aes ofb模式整合测试
import (
	"crypto/cipher"
	"crypto/aes"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func main(){
	password := "123456"
	data := []byte("hello world")

	newOfb(data, password)
	fmt.Println("data:",data);
	fmt.Println("data Hex:",hex.EncodeToString(data));

	newOfb(data, password)
	fmt.Println(string(data));
}

func Md5(password string)  []byte{
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("test md5 encrypto"))
	cipherStr := md5Ctx.Sum(nil)
	return []byte(hex.EncodeToString(cipherStr))
}

func newOfb(data []byte,key string){
	c, err := aes.NewCipher(Md5(key))
	if err != nil {
		fmt.Printf("%s: NewCipher(%d bytes) = %s \n", data, len(key), err)
	}
	iv := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	ofb := cipher.NewOFB(c, iv)
	copyTmp := data
	ofb.XORKeyStream(data, copyTmp)
}
