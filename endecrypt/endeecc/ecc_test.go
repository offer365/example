package endeecc

import (
	"encoding/hex"
	"fmt"
	"testing"
)

const (
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
)

func TestEccDecrypt(t *testing.T) {
	GetEccKey()
	var err error
	var text = []byte("hehe")
	byt,_:=EccEncrypt(text,[]byte(pub))
	fmt.Println(hex.EncodeToString(byt),err)
	byt,err=EccDecrypt(byt,[]byte(pri))
	fmt.Println(string(byt))

	// byt,err=goEncrypt.EccEncrypt(text,[]byte(pub))
	// fmt.Println(hex.EncodeToString(byt),err)
	// byt,err=goEncrypt.EccDecrypt(byt,[]byte(pri))
	// fmt.Println(string(byt))
}
