package endersa

// https://github.com/wenzhenxi/gorsa

import (
	"bytes"
	"crypto/rsa"
	"errors"
	"io/ioutil"
)

var RSA = &RSASecurity{}

type RSASecurity struct {
	pubStr []byte          //公钥字符串
	priStr []byte           //私钥字符串
	pubKey *rsa.PublicKey  //公钥
	priKey *rsa.PrivateKey //私钥
}

// 设置公钥
func (r *RSASecurity) SetPublicKey(pubStr []byte) (err error) {
	r.pubStr = pubStr
	r.pubKey, err = r.GetPublickey()
	return err
}

// 设置私钥
func (r *RSASecurity) SetPrivateKey(priStr []byte) (err error) {
	r.priStr = priStr
	r.priKey, err = r.GetPrivatekey()
	return err
}

// *rsa.PublicKey
func (r *RSASecurity) GetPrivatekey() (*rsa.PrivateKey, error) {
	return getPriKey(r.priStr)
}

// *rsa.PrivateKey
func (r *RSASecurity) GetPublickey() (*rsa.PublicKey, error) {
	return getPubKey(r.pubStr)
}

// 公钥加密
func (r *RSASecurity) PubKeyENCTYPT(input []byte) ([]byte, error) {
	if r.pubKey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(r.pubKey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 公钥解密
func (r *RSASecurity) PubKeyDECRYPT(input []byte) ([]byte, error) {
	if r.pubKey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(r.pubKey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 私钥加密
func (r *RSASecurity) PriKeyENCTYPT(input []byte) ([]byte, error) {
	if r.priKey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(r.priKey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// 私钥解密
func (r *RSASecurity) PriKeyDECRYPT(input []byte) ([]byte, error) {
	if r.priKey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(r.priKey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}

	return ioutil.ReadAll(output)
}
