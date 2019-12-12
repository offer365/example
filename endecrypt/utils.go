package endecrypt

import (
	"errors"

	"github.com/offer365/example/endecrypt/endeaes"
	"github.com/offer365/example/endecrypt/endeaesrsa"
	"github.com/offer365/example/endecrypt/endersa"
)

type method int

const (
	Pri1AesRsa1024 method = iota + 1
	Pub1AesRsa1024
	Pri1AesRsa2048
	Pub1AesRsa2048
	Pri1Rsa1024
	Pub1Rsa1024
	Pri1Rsa2048
	Pub1Rsa2048
	Aes1key16
	Aes1key32
	Pri2AesRsa1024
	Pub2AesRsa1024
	Pri2AesRsa2048
	Pub2AesRsa2048
	Pri2Rsa1024
	Pub2Rsa1024
	Pri2Rsa2048
	Pub2Rsa2048
	Aes2key16
	Aes2key32
)

func Encrypt(m method, src []byte) ([]byte, error) {
	// base64.StdEncoding.EncodeToString(byt)
	switch m {
	case Pri1AesRsa1024: // 私钥加密 1024
		return endeaesrsa.PriEncrypt(src, []byte(_pri1Key1024), []byte(_aes1Key16))
	case Pub1AesRsa1024: // 公钥加密 1024
		return endeaesrsa.PubEncrypt(src, []byte(_pub1Key1024), []byte(_aes1Key16))
	case Pri1AesRsa2048: // 私钥加密 2048
		return endeaesrsa.PriEncrypt(src, []byte(_pri1Key2048), []byte(_aes1Key32))
	case Pub1AesRsa2048: // 公钥加密 2048
		return endeaesrsa.PubEncrypt(src, []byte(_pub1Key2048), []byte(_aes1Key32))
	case Pri1Rsa1024: // 私钥加密 1024
		return endersa.PriEncrypt(src, []byte(_pri1Key1024))
	case Pub1Rsa1024: // 公钥加密 1024
		return endersa.PubEncrypt(src, []byte(_pub1Key1024))
	case Pri1Rsa2048: // 私钥加密 2048
		return endersa.PriEncrypt(src, []byte(_pri1Key2048))
	case Pub1Rsa2048: // 公钥加密 2048
		return endersa.PubEncrypt(src, []byte(_pub1Key2048))
	case Aes1key16: // aes加密 16
		return endeaes.AesCbcEncrypt(src, []byte(_aes1Key16))
	case Aes1key32: // aes加密 32
		return endeaes.AesCbcEncrypt(src, []byte(_aes1Key32))
	case Pri2AesRsa1024: // 私钥加密 1024
		return endeaesrsa.PriEncrypt(src, []byte(_pri2Key1024), []byte(_aes2Key16))
	case Pub2AesRsa1024: // 公钥加密 1024
		return endeaesrsa.PubEncrypt(src, []byte(_pub2Key1024), []byte(_aes2Key16))
	case Pri2AesRsa2048: // 私钥加密 2048
		return endeaesrsa.PriEncrypt(src, []byte(_pri2Key2048), []byte(_aes2Key32))
	case Pub2AesRsa2048: // 公钥加密 2048
		return endeaesrsa.PubEncrypt(src, []byte(_pub2Key2048), []byte(_aes2Key32))
	case Pri2Rsa1024: // 私钥加密 1024
		return endersa.PriEncrypt(src, []byte(_pri2Key1024))
	case Pub2Rsa1024: // 公钥加密 1024
		return endersa.PubEncrypt(src, []byte(_pub2Key1024))
	case Pri2Rsa2048: // 私钥加密 2048
		return endersa.PriEncrypt(src, []byte(_pri2Key2048))
	case Pub2Rsa2048: // 公钥加密 2048
		return endersa.PubEncrypt(src, []byte(_pub2Key2048))
	case Aes2key16: // aes加密 16
		return endeaes.AesCbcEncrypt(src, []byte(_aes2Key16))
	case Aes2key32: // aes加密 32
		return endeaes.AesCbcEncrypt(src, []byte(_aes2Key32))
	default:
		return nil, errors.New("method error")
	}
}

func Decrypt(m method, src []byte) ([]byte, error) {
	switch m {
	case Pri1AesRsa1024: // 私钥解密 1024
		return endeaesrsa.PriDecrypt(src, []byte(_pri1Key1024), []byte(_aes1Key16))
	case Pub1AesRsa1024: // 公钥解密 1024
		return endeaesrsa.PubDecrypt(src, []byte(_pub1Key1024), []byte(_aes1Key16))
	case Pri1AesRsa2048: // 私钥解密 2048
		return endeaesrsa.PriDecrypt(src, []byte(_pri1Key2048), []byte(_aes1Key32))
	case Pub1AesRsa2048: // 公钥解密 2048
		return endeaesrsa.PubDecrypt(src, []byte(_pub1Key2048), []byte(_aes1Key32))
	case Pri1Rsa1024: // 私钥解密 1024
		return endersa.PriDecrypt(src, []byte(_pri1Key1024))
	case Pub1Rsa1024: // 公钥解密 1024
		return endersa.PubDecrypt(src, []byte(_pub1Key1024))
	case Pri1Rsa2048: // 私钥解密 2048
		return endersa.PriDecrypt(src, []byte(_pri1Key2048))
	case Pub1Rsa2048: // 公钥解密 2048
		return endersa.PubDecrypt(src, []byte(_pub1Key2048))
	case Aes1key16: // aes解密 16
		return endeaes.AesCbcDecrypt(src, []byte(_aes1Key16))
	case Aes1key32: // aes解密 32
		return endeaes.AesCbcDecrypt(src, []byte(_aes1Key32))
	case Pri2AesRsa1024: // 私钥解密 1024
		return endeaesrsa.PriDecrypt(src, []byte(_pri2Key1024), []byte(_aes2Key16))
	case Pub2AesRsa1024: // 公钥解密 1024
		return endeaesrsa.PubDecrypt(src, []byte(_pub2Key1024), []byte(_aes2Key16))
	case Pri2AesRsa2048: // 私钥解密 2048
		return endeaesrsa.PriDecrypt(src, []byte(_pri2Key2048), []byte(_aes2Key32))
	case Pub2AesRsa2048: // 公钥解密 2048
		return endeaesrsa.PubDecrypt(src, []byte(_pub2Key2048), []byte(_aes2Key32))
	case Pri2Rsa1024: // 私钥解密 1024
		return endersa.PriDecrypt(src, []byte(_pri2Key1024))
	case Pub2Rsa1024: // 公钥解密 1024
		return endersa.PubDecrypt(src, []byte(_pub2Key1024))
	case Pri2Rsa2048: // 私钥解密 2048
		return endersa.PriDecrypt(src, []byte(_pri2Key2048))
	case Pub2Rsa2048: // 公钥解密 2048
		return endersa.PubDecrypt(src, []byte(_pub2Key2048))
	case Aes2key16: // aes解密 16
		return endeaes.AesCbcDecrypt(src, []byte(_aes2Key16))
	case Aes2key32: // aes解密 32
		return endeaes.AesCbcDecrypt(src, []byte(_aes2Key32))
	default:
		return nil, errors.New("method error")
	}
}
