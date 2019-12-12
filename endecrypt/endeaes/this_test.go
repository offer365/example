package endeaes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var (
	//指定密钥
	key = []byte("3N4QGQ8jf55KzVJaAERrp5THqVj25ikU")
	text = []byte("hello world")
)

func BenchmarkAesCbcCrypt(b *testing.B) {
	// 加密
	result, err := AesCbcEncrypt(text, key)
	if err != nil {
		panic(err)
	}
	// nRmbAgLEsFSZzieUekELhA==
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	oriData, err := AesCbcDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(oriData))
}

func BenchmarkAesCrtCrypt(b *testing.B) {
	cipherText:= AesCrtCrypt(text,key)
	fmt.Println(base64.StdEncoding.EncodeToString(cipherText))
	//解密
	plainText:= AesCrtCrypt(cipherText,key)
	fmt.Println(string(plainText))

}

func BenchmarkAesGcmCrypt(b *testing.B) {
	nonce := "5f8e40324b5d86d7483308ac"  // len = 24
	cipherText := AesGcmEncrypt(string(text), string(key), nonce)
	newTest := AesGcmDecrypt(cipherText, string(key), nonce)

	fmt.Println(string(text))
	fmt.Println(cipherText)
	fmt.Println(string(newTest))
}

func BenchmarkAesCfmCrypt(b *testing.B) {
	arrEncrypt, err := AesCfmEncrypt(text,key)
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	strMsg, err := AesCfmDecrypt(text,key)
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	fmt.Println(strMsg)
}

