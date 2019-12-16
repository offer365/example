package endedes

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

// 补码
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	// 计算需要补几位数
	padding := blockSize - len(cipherText)%blockSize
	// 在切片后面追加char数量的byte(char)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 返回填充后的密文
	return append(cipherText, padText...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadDing := int(src[length-1])
	if length < unPadDing{
		return nil
	}
	return src[:(length - unPadDing)]
}

//CBC加密
func DesCbcEncrypt(src ,key []byte)([]byte,error){
	if len(key)!=8{
		return nil,errors.New("The DES key should be,either 8 bytes.")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(src, block.BlockSize())

	//获取CBC加密模式
	iv := key //用密钥作为向量(不建议这样使用)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte,len(paddingText))
	blockMode.CryptBlocks(cipherText,paddingText)
	return cipherText,nil
}
//CBC解密
func DesCbcDecrypt(cipherText ,key []byte,ivDes...byte) ([]byte,error){
	if len(key)!=8{
		return nil,errors.New("The DES key should be,either 8 bytes.")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//获取CBC解密模式
	iv := key //用密钥作为向量(不建议这样使用)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte,len(cipherText))
	blockMode.CryptBlocks(plainText,cipherText)
	plainText = PKCS5UnPadding(plainText)
	return plainText,nil
}

func TripleDesEncrypt(src ,key []byte)([]byte,error){
	if len(key)!=24{
		return nil,errors.New("The DES key should be,either 24 bytes.")
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(src, block.BlockSize())

	//获取CBC加密模式
	iv := key[:8] //用密钥作为向量(不建议这样使用)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte,len(paddingText))
	blockMode.CryptBlocks(cipherText,paddingText)
	return cipherText,nil
}

func TripleDesDecrypt(src ,key []byte) ([]byte,error){
	if len(key)!=24{
		return nil,errors.New("The DES key should be,either 24 bytes.")
	}
	// 1. Specifies that the 3des decryption algorithm creates and returns a cipher.Block interface using the TDEA algorithm。
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	// 2. Delete the filling
	// Before deleting, prevent the user from entering different keys twice and causing panic, so do an error handling

	//获取CBC解密模式
	iv := key[:8]//用密钥作为向量(不建议这样使用)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	paddingText := make([]byte,len(src)) //
	blockMode.CryptBlocks(paddingText,src)


	plainText:= PKCS5UnPadding(paddingText)
	return plainText,nil
}

