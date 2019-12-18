package endeecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*
@Time : 2018/11/4 16:22 
@Author : wuman
@File : GetECCKey
@Software: GoLand
*/

const (
	eccPrivateKeyPrefix="ECC PRIVATE KEY "
	eccPublicKeyPrefix="ECC PUBLIC KEY "
)

func GetEccKey(eccPrivateFileName,eccPublishFileName string)error{
	// elliptic.P256() elliptic.P521()
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err!=nil{
		return err
	}

	x509PrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err!=nil{
		return err
	}

	block := pem.Block{
		Type:  eccPrivateKeyPrefix,
		Bytes: x509PrivateKey,
	}
	file, err := os.Create(eccPrivateFileName)
	if err!=nil{
		return err
	}
	defer file.Close()
	if err=pem.Encode(file, &block);err!=nil{
		return err
	}

	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err!=nil {
		return err
	}
	publicBlock := pem.Block{
		Type:  eccPublicKeyPrefix,
		Bytes: x509PublicKey,
	}
	publicFile, err := os.Create(eccPublishFileName)
	if err!=nil {
		return err
	}
	defer publicFile.Close()
	if err=pem.Encode(publicFile,&publicBlock);err!=nil{
		return err
	}
	return nil
}