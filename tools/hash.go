package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"

	"github.com/spaolacci/murmur3"
)

func Md5sum(byt []byte, salt []byte) string {
	h := md5.New()
	if salt != nil {
		byt = append(byt, salt...)
	}
	h.Write(byt)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Sha256sum(byt []byte, salt []byte) string {
	h := sha256.New()
	if salt != nil {
		byt = append(byt, salt...)
	}
	h.Write(byt)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
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
