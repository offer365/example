package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/spaolacci/murmur3"
)

func Md5Hex(byt []byte, salt []byte) string {
	h := md5.New()
	if salt != nil {
		byt = append(byt, salt...)
	}
	h.Write(byt)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256Hex(byt []byte, salt []byte) string {
	h := sha256.New()
	if salt != nil {
		byt = append(byt, salt...)
	}
	h.Write(byt)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha512Hex(byt []byte, salt []byte) string {
	h := sha512.New()
	if salt != nil {
		byt = append(byt, salt...)
	}
	h.Write(byt)
	return hex.EncodeToString(h.Sum(nil))
}

// BenchmarkMd5Hash-4               3500572               326 ns/op
// BenchmarkSha1Hash-4              2593280               423 ns/op
// BenchmarkMurmurHash32-4         80041621                14.6 ns/op
// BenchmarkMurmurHash64-4         37517820                35.3 ns/op
// murmurhash相比其它的算法有三倍以上的性能提升
func Md5Hash(byt []byte) string {
	res := md5.Sum(byt) // [16]byte
	return hex.EncodeToString(res[:])
}

func Sha1Hash(byt []byte) string {
	res := sha1.Sum(byt) // [20]byte
	return hex.EncodeToString(res[:])
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
