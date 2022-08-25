package crypto_test

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
	"testing"
)

const (
	PW_SALT_BYTES = 8
	PW_HASH_BYTES = 32
	PASS_WORD     = "hello scrypt"
)

func TestScrypt(t *testing.T) {
	// 生成8位密码盐值
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	// 加密
	// func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
	// password []byte: 密码, 字节数组表示
	// salt []byte: 密码盐值, 字节数组表示
	// N int: CPU的核心数, 必须是2幂级数并且大于1
	// r, p int: r, p须满足r * p < 2³⁰, 推荐 r=8, p=1
	// keyLen int: 密码字节数组的长度
	hash, err := scrypt.Key([]byte(PASS_WORD), salt, 1<<15, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("len(salt)-----------------------", len(salt))
	log.Println("len(hash)-----------------------", len(hash))

	passSalt := base64.StdEncoding.EncodeToString(salt)
	hashPass := base64.StdEncoding.EncodeToString(hash)

	log.Println("passSalt-------------------------", passSalt)
	log.Println("hashPass-------------------------", hashPass)
}

