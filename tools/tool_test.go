package tools

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	fmt.Println(Abs(-7458))
	fmt.Println(Abs2(-7458))
}

func TestRandString(t *testing.T) {
	fmt.Println(RandString(32))
	fmt.Println(RandString(32))

}

func TestRunTime(t *testing.T) {
	fmt.Println(RunTime(time.Now().Unix(), time.Now().Unix()-1234567))
}

func TestMd5sum(t *testing.T) {
	str := RandString(16)
	fmt.Println(str)
	fmt.Println(Md5sum([]byte(str), []byte("")))
	fmt.Println(Sha256sum([]byte(str), []byte("")))

}

func TestMd5sum2(t *testing.T) {
	var byt = []byte("hello world")
	fmt.Println(Md5Hash(byt))
	fmt.Println(Sha1Hash(byt))
	fmt.Println(Murmur32(byt))
	fmt.Println(Murmur64(byt))
	fmt.Println(Murmur128(byt))
}

func TestScrypt(t *testing.T) {
	str, err := Scrypt([]byte("hehe"), []byte("123"))
	fmt.Println(base64.StdEncoding.EncodeToString(str), err)
}

func TestGenerateFromPassword(t *testing.T) {
	cipher, err := GenerateFromPassword([]byte("123456"))
	fmt.Println(string(cipher), err)
	err = CompareHashAndPassword(cipher, []byte("123456"))
	if err == nil {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	err = CompareHashAndPassword(cipher, []byte("1234567"))
	if err == nil {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
