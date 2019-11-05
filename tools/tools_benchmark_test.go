package tools

import "testing"


// go test -test.bench=.
var byt=[]byte("hello world")

func BenchmarkMd5Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5Hash(byt)
	}
}

func BenchmarkSha1Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha1Hash(byt)
	}
}

func BenchmarkMurmurHash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur32(byt)
	}
}

func BenchmarkMurmurHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur64(byt)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(50-100)
	}
}

func BenchmarkAbs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs2(50-100)
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(32)
	}
}

func BenchmarkRandStringRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes(32)
	}
}

func BenchmarkMurmur128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur128(byt)
	}
}

func BenchmarkScrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Scrypt([]byte("hehehehehehehehehehehehehehehehehehehehehehe"),[]byte("123123123123123"))
	}
}
