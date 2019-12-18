package endeaesecc

import (
	"github.com/offer365/example/endecrypt/endeaes"
	"github.com/offer365/example/endecrypt/endeecc"
)

// 公钥加密
func PubEncrypt(src,eccKey, aesKey []byte) (byt []byte, err error) {
	// RSA公钥加密
	byt, err = endeecc.EccEncrypt(src, eccKey)
	if err != nil {
		return
	}
	// AES加密
	return endeaes.AesCbcEncrypt(byt, aesKey)
	//return base64.StdEncoding.EncodeToString(result), err
}


// 私钥解密
func PriDecrypt(src, eccKey, aesKey []byte) (byt []byte,err error) {
	//byt, err := base64.StdEncoding.DecodeString(src)
	// aes 解密
	byt, err = endeaes.AesCbcDecrypt(src, aesKey)
	if err != nil {
		return
	}
	// rsa 私钥解密
	return endeecc.EccDecrypt(byt, eccKey)
}