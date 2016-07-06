package main
//rc4流加密测试
import(
	"crypto/rc4"
	"fmt"
	"encoding/base64"
)

func main(){
	src := "hello world！！！！！！！！！！！！！！123";
	srcByte := []byte(src)
	dst := make([]byte,len(src))
	key := []byte("asfasasdf;")

	//加密
	crypt(srcByte,dst,key);

	fmt.Println("src(",len(srcByte),"):",srcByte)
	fmt.Println("src(",len(dst),"):",dst)

	fmt.Println(base64.StdEncoding.EncodeToString(dst));

	//解密
	crypt(dst,dst,key);

	fmt.Println(string(dst));
}

func crypt(src, dst, key []byte)(err error){
	c, err := rc4.NewCipher(key)
	c.XORKeyStream(dst, src)
	return
}