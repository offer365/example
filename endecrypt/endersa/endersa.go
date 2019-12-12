package endersa

// https://github.com/wenzhenxi/gorsa

// 公钥加密
func PubEncrypt(src, publicKey []byte) ([]byte, error) {

	gRsa := RSASecurity{}
	gRsa.SetPublicKey(publicKey)

	rsaData, err := gRsa.PubKeyENCTYPT(src)
	if err != nil {
		return nil, err
	}
	// base64.StdEncoding.EncodeToString(rsaData), nil
	return rsaData, nil
}

// 私钥加密
func PriEncrypt(src, privateKey []byte) ([]byte, error) {

	gRsa := RSASecurity{}
	gRsa.SetPrivateKey(privateKey)

	rsaData, err := gRsa.PriKeyENCTYPT(src)
	if err != nil {
		return nil, err
	}

	// base64.StdEncoding.EncodeToString(rsaData), nil
	return rsaData, nil
}

// 公钥解密
func PubDecrypt(src, publicKey []byte) ([]byte, error) {

	//dataByt, _ := base64.StdEncoding.DecodeString(src)

	gRsa := RSASecurity{}
	gRsa.SetPublicKey(publicKey)

	rsaData, err := gRsa.PubKeyDECRYPT(src)
	if err != nil {
		return nil, err
	}

	return rsaData, nil

}

// 私钥解密
func PriDecrypt(src, privateKey []byte) ([]byte, error) {

	//dataByt, _ := base64.StdEncoding.DecodeString(src)

	gRsa := RSASecurity{}
	gRsa.SetPrivateKey(privateKey)

	rsaData, err := gRsa.PriKeyDECRYPT(src)
	if err != nil {
		return nil, err
	}

	return rsaData, nil
}
