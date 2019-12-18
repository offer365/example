package endeaesrsaecc

import (
	"github.com/offer365/example/endecrypt/endeaes"
	"github.com/offer365/example/endecrypt/endeecc"
	"github.com/offer365/example/endecrypt/endersa"
)

// 公钥加密
func PubEncrypt(src,ecckey,rsaKey, aesKey []byte) (byt []byte, err error) {
	byt, err = endeecc.EccEncrypt(src,ecckey)
	// RSA公钥加密
	byt, err = endersa.PubEncrypt(byt, rsaKey)
	if err != nil {
		return
	}
	// AES加密
	return endeaes.AesCbcEncrypt(byt, aesKey)
	//return base64.StdEncoding.EncodeToString(result), err
}

// 私钥解密
func PriDecrypt(src, ecckey,rsaKey, aesKey []byte) (byt []byte,err error) {
	//byt, err := base64.StdEncoding.DecodeString(src)
	// aes 解密
	byt, err = endeaes.AesCbcDecrypt(src, aesKey)
	if err != nil {
		return
	}
	// rsa 私钥解密
	byt, err =  endersa.PriDecrypt(byt, rsaKey)
	return endeecc.EccDecrypt(byt,ecckey)
}