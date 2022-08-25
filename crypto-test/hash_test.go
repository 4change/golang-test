package crypto_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"testing"
)

func Test_MD5(t *testing.T) {
	testString := "Hi,Spider Man !"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(testString))
	Result := Md5Inst.Sum([]byte(""))

	fmt.Printf("%x\n\n", Result)
}

func Test_SHA1(t *testing.T) {
	TestString := "Hi,Spider Man !"

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result := Sha1Inst.Sum([]byte(""))

	fmt.Printf("%x\n\n", Result)
}

func Test_SHA256(t *testing.T) {
	TestString := "Hi,Spider Man !"

	Sha256Inst := sha256.New()
	Sha256Inst.Write([]byte(TestString))
	Result := Sha256Inst.Sum([]byte(""))

	base64OfSha256 := base64.StdEncoding.EncodeToString(Result)

	fmt.Println(base64OfSha256)
}

func Test_Hash_Salt(t *testing.T) {

	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "需要加密的密码")

	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last :=fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(last)
}
