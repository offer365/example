package endedes

import (
	"fmt"
	"testing"
)

func TestCbc(t *testing.T) {
	//key的长度必须都是8位
	var key = []byte("12345678")
	var text = []byte("hehe")

	byt,_:= DesCbcEncrypt(text,key)
	// fmt.Println(string(byt))
	byt,_ = DesCbcDecrypt(byt, key)
	fmt.Println(string(byt))
	byt,_=TripleDesEncrypt(text,[]byte("123456781234567812345678"))
	byt,_ = TripleDesDecrypt(byt, []byte("123456781234567812345678"))
	fmt.Println(string(byt))
}