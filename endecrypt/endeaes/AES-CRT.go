package endeaes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

//AEC加密和解密（CRT模式）
func AesCrtCrypt(src,k []byte) ([]byte,error){
	length := len(k)
	if length%8 != 0 || length < 16 || length > 32 {
		return nil,errors.New("The AES key should be,either 16, 24, or 32 bytes.")
	}
	//指定加密、解密算法为AES，返回一个AES的Block接口对象
	block,err:=aes.NewCipher(k)
	if err!=nil{
		return nil,err
	}
	//指定计数器,长度必须等于block的块尺寸
	count:=[]byte("j25zVJa3N4QGJa3NQ8jf55KzVJaAERrp5THqVikU")[:block.BlockSize()]
	//指定分组模式
	blockMode:=cipher.NewCTR(block,count)
	//执行加密、解密操作
	cipherText:=make([]byte,len(src))
	blockMode.XORKeyStream(cipherText, src)
	//返回明文或密文
	return cipherText,nil
}

