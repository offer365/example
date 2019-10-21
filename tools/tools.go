package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

// 随机生成字符串指定个数的字符串
func RandString(n int) string {
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 运行时间
func RunTime(now, start int64) string {
	online := now - start
	d := online / 86400
	h := (online - d*86400) / 3600
	m := (online - d*86400 - h*3600) / 60
	s := online - d*86400 - h*3600 - m*60
	return fmt.Sprintf("%02d天%02d小时%02d分钟%02d秒.", d, h, m, s)
}

// 绝对值
func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func Md5sum(byt []byte,salt []byte) string  {
	h := md5.New()
	h.Write(byt)
	return base64.StdEncoding.EncodeToString(h.Sum(salt))
}

func Sha256sum(byt []byte,salt []byte) string {
	h:=sha256.New()
	h.Write(byt)
	return base64.StdEncoding.EncodeToString(h.Sum(salt))
}

