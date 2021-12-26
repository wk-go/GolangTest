package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

//自定义Claims
type CustomClaims struct {
	UserId             int64 `json:"user_id"`
	jwt.StandardClaims `json:"standard_claims"`
}

func main() {
	//生成token
	maxAge := 60 * 60 * 24
	customClaims := &CustomClaims{
		UserId: 11, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    "jerry",                                                    // 非必须，也可以填充用户名，
		},
	}
	token, err := CreateToken(customClaims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token:", token)

	//解析token
	ret, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("userinfo: %+v\n", ret)
}

func CreateToken(data jwt.Claims) (token string, err error) {
	//采用HMAC SHA256加密算法
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err = _token.SignedString([]byte(SECRETKEY))
	return
}

//解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
