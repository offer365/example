package endeaes

import (
	"crypto/aes"
	"crypto/cipher"
)

//加密字符串
func AesCfmEncrypt(src,k []byte) ([]byte, error) {
	length := len(k)
	if length%8 != 0 || length < 16 || length > 32 {
		panic("The AES key should be,either 16, 24, or 32 bytes.")
	}
	iv:=k[:aes.BlockSize]
	cipherText := make([]byte, len(src))
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	aesCfb:= cipher.NewCFBEncrypter(block, iv)
	aesCfb.XORKeyStream(cipherText, src)
	return cipherText, nil
}

//解密字符串
func AesCfmDecrypt(src,k []byte) (string,error) {
	length := len(k)
	if length%8 != 0 || length < 16 || length > 32 {
		panic("The AES key should be,either 16, 24, or 32 bytes.")
	}
	iv:=k[:aes.BlockSize]
	plainText := make([]byte, len(src))
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	aesCfb := cipher.NewCFBDecrypter(block, iv)
	aesCfb.XORKeyStream(plainText, src)
	return string(plainText), nil
}
