package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spaolacci/murmur3"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"

	"math/rand"
	"time"
)

// 随机生成字符串指定个数的字符串

// BenchmarkRandString-4              81678             14439 ns/op
// BenchmarkRandStringRunes-4         71467             15565 ns/op

func RandString(n int) string {
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
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

// BenchmarkAbs-4                  1000000000               0.510 ns/op
// BenchmarkAbs2-4                 1000000000               0.460 ns/op

// 绝对值
func Abs(a int64) int64 {
	return (a ^ a>>31) - a>>31
}

// 绝对值
func Abs2(n int64) int64 {
	return (n ^ n>>63) - n>>63
}

func Md5sum(byt []byte, salt []byte) string {
	h := md5.New()
	h.Write(byt)
	return base64.StdEncoding.EncodeToString(h.Sum(salt))
}

func Sha256sum(byt []byte, salt []byte) string {
	h := sha256.New()
	h.Write(byt)
	return base64.StdEncoding.EncodeToString(h.Sum(salt))
}

// BenchmarkMd5Hash-4               3500572               326 ns/op
// BenchmarkSha1Hash-4              2593280               423 ns/op
// BenchmarkMurmurHash32-4         80041621                14.6 ns/op
// BenchmarkMurmurHash64-4         37517820                35.3 ns/op
// murmurhash相比其它的算法有三倍以上的性能提升
func Md5Hash(byt []byte) string {
	res := md5.Sum(byt) // [16]byte
	return base64.StdEncoding.EncodeToString(res[:])
}

func Sha1Hash(byt []byte) string {
	res := sha1.Sum(byt) // [20]byte
	return base64.StdEncoding.EncodeToString(res[:])
}

func Murmur32(byt []byte) uint32 {
	return murmur3.Sum32(byt)
}

func Murmur64(byt []byte) uint64 {
	return murmur3.Sum64(byt)
}

func Murmur128(byt []byte) (uint64, uint64) {
	return murmur3.Sum128(byt)
}

// 可以获取唯一的相应的密码值，这是目前为止最难破解的，但耗时比较久100ms
func Scrypt(pwd []byte, salt []byte, ) ([]byte, error) {
	return scrypt.Key(pwd, salt, 1<<15, 8, 1, 32)
}

// 在数据库中存储密码与校验密码

// 用密码生产一个密文
func GenerateFromPassword(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, 12)
}

// 将密文与密码比较。error==nil 则相等。
func CompareHashAndPassword(cipher, pwd []byte) error {
	return bcrypt.CompareHashAndPassword(cipher, pwd)
}

// 获取运行的根目录
func GetRootDir() (p string) {
	return filepath.Dir(os.Args[0])
}
