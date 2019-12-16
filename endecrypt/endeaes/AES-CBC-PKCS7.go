package endeaes

/*
AES256  CBC模式
CBC模式+PKCS7 填充方式实现AES的加密和解密
*/

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// 补码
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	// 计算需要补几位数
	padding := blockSize - len(cipherText)%blockSize
	// 在切片后面追加char数量的byte(char)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 返回填充后的密文
	return append(cipherText, padText...)
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unPadDing := int(src[length-1])
	if length < unPadDing{
		return nil
	}
	return src[:(length - unPadDing)]
}

// 加密
func AesCbcEncrypt(src, key []byte) ([]byte, error) {
	length := len(key)
	if length%8 != 0 || length < 16 || length > 32 {
		return nil,errors.New("The AES key should be,either 16, 24, or 32 bytes.")
	}
	// NewCipher创建并返回一个新的cipher.Block。 关键参数应该是AES密钥， 要选择16个，24个或32个字节 AES-128，AES-192或AES-256。
	// 获取block块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	// 补码
	src = PKCS7Padding(src, blockSize)
	// NewCBCEncrypter返回一个BlockMode，它以密码块链接加密 模式，使用给定的块。 iv的长度必须与 块的块大小。
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 创建明文长度的数组
	cipherText := make([]byte, len(src))
	// 加密明文
	blockMode.CryptBlocks(cipherText, src)
	return cipherText, nil
}

// 解密
func AesCbcDecrypt(src, key []byte) ([]byte, error) {
	length := len(key)
	if length%8 != 0 || length < 16 || length > 32 {
		return nil,errors.New("The AES key should be,either 16, 24, or 32 bytes.")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	plainText := make([]byte, len(src))
	blockMode.CryptBlocks(plainText, src)
	plainText = PKCS7UnPadding(plainText)
	return plainText, nil
}
