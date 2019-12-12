package endeaesrsa

import (
	"encoding/base64"
	"fmt"

	"testing"
)
var (
	//指定密钥
	text = []byte("hello world")
	err error
	str string
)


// 私钥加密
func TestPriEncrypt(t *testing.T) {
	byt,err:=PriEncrypt(text,[]byte(_pri1Key2048),[]byte(_aes1Key32))
	base64.StdEncoding.EncodeToString(byt)
	tt,err:=PubDecrypt(byt,[]byte(_pub1Key2048),[]byte(_aes1Key32))
	fmt.Println(err)
	fmt.Println(string(tt))
}

// 公钥加密
func TestPubEncrypt(t *testing.T) {
	byt,err:=PubEncrypt(text,[]byte(_pub2Key2048),[]byte(_aes2Key32))
	base64.StdEncoding.EncodeToString(byt)
	tt,err:=PriDecrypt(byt,[]byte(_pri2Key2048),[]byte(_aes2Key32))
	fmt.Println(err)
	fmt.Println(string(tt))
}