package main

import (
	"encoding/pem"
	"errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"fmt"
	"encoding/base64"
)

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDkAExKmo9ka8sfKw2bABG935xV30NJrXHBQ/qKcXtsMcofpE3+
3WNrH1pjD2NRlOZGVmvhAhGp0KiMZ6k5bWF5rukGPpzEsazbo1y73ldEOjru60n4
0P7N1DuoteYWdk0KlQb9lt2HBtpm/uWxOQzImOvMZ51BnsvDjF0nVBiqnQIDAQAB
AoGABey4Dsw7Y6mlapbs0JVM4Lk5z8Vwcy6toQ8KKKTQRzx3+yCC4leQaM00xRQ2
SX1sCnHedcde/CGu748WB6b+/GTshcj2v5WU9NlFTcqGkIKgu/WnmPwPau0sMvU6
fQR7r6jvAI3euXIPphrXRJ3gajU7D4CBlrd/wn/i4rh0KgUCQQD8yWyZ35kKj4/t
P9D5g8iM7yH5zD7kGvcbv3kyT6Y5neRf4m2Sj/jXtKiAWFI8Hqj758av5iIw9/CT
J2mpq49HAkEA5uY4djV/LsZYJ278DmmAAeQjxYNXqmMDlBaMMvKMpwHOLO0FIzpE
QexyNPz+AglRuWLqkwXbxT5nPsBl03tQ+wJAMqENFUyJVGoog3YSnsbcNg33Ghbk
Sb902qPg3EjDnCqZgPLSy1X2mw1d6kbGQbBKXBmx260WEAS4tGBic08fJQJAXgX7
ke9A5gwwk4Y3L6s4TAzZoDFWvnRpXaE83/Yy3kL28QZnZCvy5aFh9D/dM3kWBVbJ
TKtDDfPWWRBBprd9hwJBAMgOauq4V+AyFJHe5Unn8+832xKoPTRQbn+fiWdxqIdn
giRXeBMVMFA+rl3oiSy1CSXhFAvRLNCDDt74LT+X9rg=
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDkAExKmo9ka8sfKw2bABG935xV
30NJrXHBQ/qKcXtsMcofpE3+3WNrH1pjD2NRlOZGVmvhAhGp0KiMZ6k5bWF5rukG
PpzEsazbo1y73ldEOjru60n40P7N1DuoteYWdk0KlQb9lt2HBtpm/uWxOQzImOvM
Z51BnsvDjF0nVBiqnQIDAQAB
-----END PUBLIC KEY-----
`)

// 公钥加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 私钥解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func main() {
	data := "hello world!!!,hello china!!!"
	en_data, _ := RsaEncrypt([]byte(data))
	en_data_base64 := base64.StdEncoding.EncodeToString(en_data)
	fmt.Println(en_data_base64)
	de_data_base64,_ := base64.StdEncoding.DecodeString(en_data_base64)
	de_data, _ := RsaDecrypt(de_data_base64)
	fmt.Println(string(de_data))
}
