package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	fmt.Println(Abs(-7458))
}

func TestRandString(t *testing.T) {
	fmt.Println(RandString(32))
}

func TestRunTime(t *testing.T) {
	fmt.Println(RunTime(time.Now().Unix(),time.Now().Unix()-1234567))
}

func TestMd5sum(t *testing.T) {
	for range time.Tick(10*time.Second){
		str:=RandString(16)
		fmt.Println("......................................")
		fmt.Println(str)
		fmt.Println(Md5sum([]byte(str),[]byte("")))
		fmt.Println(Sha256sum([]byte(str),[]byte("")))
	}
}

