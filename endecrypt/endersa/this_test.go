package endersa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var (
	text = []byte("hello world.")
)

// 私钥加密
func TestPriEncrypt(t *testing.T) {
	// 加密
	result, err := PriEncrypt(text, []byte(_pri1Key1024))
	if err != nil {
		panic(err)
	}
	// nRmbAgLEsFSZzieUekELhA==
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	oriData, err := PubDecrypt(result, []byte(_pub1Key1024))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(oriData))
}

// 公钥加密
func TestPubEncrypt(t *testing.T) {
	// 加密
	result, err := PubEncrypt(text, []byte(_pub2Key2048))
	if err != nil {
		panic(err)
	}
	// nRmbAgLEsFSZzieUekELhA==
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	oriData, err := PriDecrypt(result, []byte(_pri2Key2048))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(oriData))
}
