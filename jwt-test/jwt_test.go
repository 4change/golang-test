package jwt_test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

// 自定义的签名
var SecretKey  =  []byte("mySigningKey")

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid uint `json:"uid"`
	Admin bool `json:"admin"`
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(SecretKey []byte, issuer string, Uid uint, isAdmin bool) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),		// 签名过期时间
			Issuer:    issuer,												// 签名发行者
		},
		Uid,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

func TestCreateToken(t *testing.T)  {
	token, _ := CreateToken([]byte(SecretKey), "root", 1, true)
	fmt.Println(token)
}

/**
 * 解析 token
 *		使用自定义的SecretKey解析出token
 */
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}

func TestCreateToken2(t *testing.T)  {
	token, _ := CreateToken([]byte(SecretKey), "root", 1, true)
	fmt.Println(token)

	claims, err := ParseToken(token, []byte(SecretKey))
	if nil != err {
		fmt.Println(" err :", err)
	}
	fmt.Println("claims:", claims)
	fmt.Println("claims uid:", claims.(jwt.MapClaims)["uid"])
}