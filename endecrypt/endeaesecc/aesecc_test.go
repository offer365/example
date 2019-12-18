package endeaesecc

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const  (
	pri=`
-----BEGIN ECC PRIVATE KEY -----
MHcCAQEEIFcWCOLVJ2VoDzKSZiNXUcVOqhlt2i9/k9urCCgCJ4TaoAoGCCqGSM49
AwEHoUQDQgAEeRcuzZ6fPlixH02gJG5c3laWMxWySeD/JBPL6fbSgj3YPl8x3AYm
bHDrAnpe1BAMZPbAARuojZTAkCDhp7TTkA==
-----END ECC PRIVATE KEY -----
`
	pub=`
-----BEGIN ECC PUBLIC KEY -----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEeRcuzZ6fPlixH02gJG5c3laWMxWy
SeD/JBPL6fbSgj3YPl8x3AYmbHDrAnpe1BAMZPbAARuojZTAkCDhp7TTkA==
-----END ECC PUBLIC KEY -----
`
	aes=`f7e8b819l0ad0ccf9a9g8fc5e8c4765q`
)

func TestPubEncrypt(t *testing.T) {
	text:=[]byte("hello world.")
	byt,err:=PubEncrypt(text,[]byte(pub),[]byte(aes))
	fmt.Println(base64.StdEncoding.EncodeToString(byt),err)
	byt,err=PriDecrypt(byt,[]byte(pri),[]byte(aes))
	fmt.Println(string(byt),err)
}
