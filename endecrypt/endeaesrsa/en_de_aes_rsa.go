package endeaesrsa

import (
	"github.com/offer365/example/endecrypt/endeaes"
	"github.com/offer365/example/endecrypt/endersa"
)

// 公钥加密
func PubEncrypt(src,rsaKey, aesKey []byte) (byt []byte, err error) {
	// RSA公钥加密
	byt, err = endersa.PubEncrypt(src, rsaKey)
	if err != nil {
		return
	}
	// AES加密
	return endeaes.AesCbcEncrypt(byt, aesKey)
	//return base64.StdEncoding.EncodeToString(result), err
}

// 公钥解密
func PubDecrypt(src, rsaKey, aesKey []byte) (byt []byte,err error) {
	//byt, err := base64.StdEncoding.DecodeString(src)
	// aes 解密
	byt, err = endeaes.AesCbcDecrypt(src, aesKey)
	if err != nil {
		return
	}
	// rsa 公钥解密
	return endersa.PubDecrypt(byt, rsaKey)
}

// 私钥加密
func PriEncrypt(src,rsaKey, aesKey []byte) (byt []byte,err error) {
	byt, err = endersa.PriEncrypt(src, rsaKey)
	if err != nil {
		return
	}
	// AES加密
	return endeaes.AesCbcEncrypt(byt, aesKey)
	//return base64.StdEncoding.EncodeToString(result), nil
}

// 私钥解密
func PriDecrypt(src, rsaKey, aesKey []byte) (byt []byte,err error) {
	//byt, err := base64.StdEncoding.DecodeString(src)
	// aes 解密
	byt, err = endeaes.AesCbcDecrypt(src, aesKey)
	if err != nil {
		return
	}
	// rsa 私钥解密
	return endersa.PriDecrypt(byt, rsaKey)
}

