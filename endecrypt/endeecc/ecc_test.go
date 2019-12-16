package endeecc

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/wumansgy/goEncrypt"
)

const (
	pri=`
-----BEGIN  ECC PRIVATE KEY -----
MHcCAQEEIJB+rphsibKVfra9h8XDnTZvmPx3eFAhWlQhV6aLja/NoAoGCCqGSM49
AwEHoUQDQgAEJWLY/c92a9sULfndiFzdOSN/d3HKX2g+vyjjaC/R2tB0dvwNK62w
fQusekv8IkyVTZI1kdytfAibXcj2sRMPTA==
-----END  ECC PRIVATE KEY -----
`
	pub=`
-----BEGIN  ECC PUBLIC KEY -----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJWLY/c92a9sULfndiFzdOSN/d3HK
X2g+vyjjaC/R2tB0dvwNK62wfQusekv8IkyVTZI1kdytfAibXcj2sRMPTA==
-----END  ECC PUBLIC KEY -----
`
)

func TestEccDecrypt(t *testing.T) {
	// GetEccKey()
	var text = []byte("hehe")
	byt,err:=EccEncrypt(text,[]byte(pub))
	fmt.Println(hex.EncodeToString(byt),err)
	byt,err=EccDecrypt(byt,[]byte(pri))
	fmt.Println(string(byt))

	byt,err=goEncrypt.EccEncrypt(text,[]byte(pub))
	fmt.Println(hex.EncodeToString(byt),err)
	byt,err=goEncrypt.EccDecrypt(byt,[]byte(pri))
	fmt.Println(string(byt))
}
