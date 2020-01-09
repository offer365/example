package endersa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var (
	text = []byte("hello world.")
)


const (
	pub1 =`
-----BEGIN RSA PUBLIC KEY-----
MIIBCAKCAQEA2WnRfmi3wCy3xliWF1FlOFusXexrGfYlu2oVC/UoZoooUvG6ovia
KkezQaw1gyOz7JMsSF8sH1V4wpV/xdCXMzsugiyPGPaFQvr8iErKj0Ga8QHPCW+5
UQadM8zS4oA0HhPsmyr0/ZGB2wiGNX8YXXCZYyYwVdeB94t0c+mb+IqFL6rGJG7u
YVt7c/NPPx2tIPTsL5QjWD8JcYdJSoYhuF3DPchHfo0spu0oAtV0g+pFqQLMaD30
izqQ8UtDj0krgLgXL16opffqjTt4mCWnQS8YUevSVIbOs4MKZuvKEHVtUoIaN2/V
kztuq+S0RVLOQQjNILkAF427q6NHdZAbDQIBAw==
-----END RSA PUBLIC KEY-----
`
	pri1 =`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA2WnRfmi3wCy3xliWF1FlOFusXexrGfYlu2oVC/UoZoooUvG6
oviaKkezQaw1gyOz7JMsSF8sH1V4wpV/xdCXMzsugiyPGPaFQvr8iErKj0Ga8QHP
CW+5UQadM8zS4oA0HhPsmyr0/ZGB2wiGNX8YXXCZYyYwVdeB94t0c+mb+IqFL6rG
JG7uYVt7c/NPPx2tIPTsL5QjWD8JcYdJSoYhuF3DPchHfo0spu0oAtV0g+pFqQLM
aD30izqQ8UtDj0krgLgXL16opffqjTt4mCWnQS8YUevSVIbOs4MKZuvKEHVtUoIa
N2/Vkztuq+S0RVLOQQjNILkAF427q6NHdZAbDQIBAwKCAQEAkPE2VEXP1XMlLuW5
ZODuJZJy6UhHZqQZJ5wOB/jFmbFwN0vRwfsRcYUiK8gjrMJ38wzIMD9yv45Qgbj/
2TW6Iid0VshfZfmuLKdTBYcxtNZnS1aKBkp7i1m+IoiMlwAivrfzEhyjU7ZWkgWu
zlS66PW7l27K4+UBT7JNopu9UFseOSA5SQkzBkJV8cSrFsKOZgPsJtfw5QyJfdnV
iCXjgTdw9It1C4F1wTDbNqLrwySMc5CauFEk+p0uAfK4XS8c+sgCBnPDc1onK5ZX
nQTUBUlMrpfR0oQx3g+fPrwfrZeJIyEK7KBOAGHyvm1jySWf0Ub93H+PpLba0+Jj
Zxxx6wKBgQDtslN+43EjQKvM3qUHeBLJMy35u8ezJZh5/YS4ZpsteJpCd7ckFqXa
BNkU7iijNdXBr4eejBE4CLLXKPRtqgKs+Ci+migWv5H0ox01GwxkgSjuJy+CTf28
JuTnV2lyr7+QQtL2TfR+dZfLlLUixUq6KNHf4GI92/f56Zoby3k+VQKBgQDqJ6bx
U2/+l0wtqqdHJOcOVMEQOYiG2xPBNzvQl7Ie/fBXVz/z1qUiAErKQrhvqV2xTCJF
x7MFCpv0xWrBWX/TEGNVi4jsuV67KLzAEZIEuBg3JNiVSsLHv4a0MWgn3FKPWv2T
horiHRC2+Yt70k+kXkxwdZdqxIN5hDWWj2wx2QKBgQCeduJUl6DCKx0zPxivpWHb
d3P70oUiGRBRU63QRGdzpbwsT89tZG6RWJC4nsXCI+PWdQUUXWDQBcyPcKLzxqxz
UBspvBq51Qv4bL4jZ12YVhtJb3UBiVPSxJiaOkZMdSpgLIykM/hUTmUyYyNsg4cm
xeE/6uwpPU/78RFn3Pt+4wKBgQCcGm9LjPVUZN1zxxovbe9e4ytgJlsEkg0reifg
ZSFp/qA6OiqijxjBVYcxgdBKcOkg3WwuhSIDXGf4g5yA5lU3YEI5B7Cd0OnSGyiA
C7at0BAkwzsOMdcv1QR4IPAakuG051O3rwdBaLXPUQen4YptlDL1o7pHLaz7rXkP
CkghOwKBgQCbO64LXg7LY/2I+QpxgkwPfyYozbCZtiLSfpZZYc+w+Dks0opAbrl6
EnKExwzBTD/gGm+eYENElj3oS3g0h5TANA5IN+hKRi9PD0UNfkx+Pq3/9wOK6i5J
TlA0FgknQHOs4ZLBhDaN2RIVcLWc6ukw52w6Zg1Jx822ONKxC3b+lg==
-----END RSA PRIVATE KEY-----
`

	pub2=`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1AFFvEgfRwJsIpiJuQxG
z7ABqW9EwreBtfk8W2yP40qqRz6xReKIv87DttKWSdKi6BvWc9T/QUhS4N+Z7iDy
ifZSkcv+e7zvtXGbe0hQxD+UlxjjhgsIBCRYL6vjV3Pn5FnxWt5h3cCeYKpTGlKv
IdDCdb1KUVlplNMTMDPkpgwoWPaBgjOV4Tv9vp+qFPRO6h/eAhOE58vGy8VVhdBm
z4+6/wDUojtckwBBaqV/FmSl1JY/iTYeAageFTGTFl1QY0Dxv/BA2GkqZktoz29w
hySmIMVnanO4GW7jyOwiBdKlucxujyOYUNsKA556SopM1iC+Z7jEKQhnWwptoiy8
+QIDAQAB
-----END PUBLIC KEY-----`
	pri2=`
-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDUAUW8SB9HAmwi
mIm5DEbPsAGpb0TCt4G1+TxbbI/jSqpHPrFF4oi/zsO20pZJ0qLoG9Zz1P9BSFLg
35nuIPKJ9lKRy/57vO+1cZt7SFDEP5SXGOOGCwgEJFgvq+NXc+fkWfFa3mHdwJ5g
qlMaUq8h0MJ1vUpRWWmU0xMwM+SmDChY9oGCM5XhO/2+n6oU9E7qH94CE4Tny8bL
xVWF0GbPj7r/ANSiO1yTAEFqpX8WZKXUlj+JNh4BqB4VMZMWXVBjQPG/8EDYaSpm
S2jPb3CHJKYgxWdqc7gZbuPI7CIF0qW5zG6PI5hQ2woDnnpKikzWIL5nuMQpCGdb
Cm2iLLz5AgMBAAECggEBAM1qXjNMfmHDSxtKSUdvSb06gKG3GhzAyYGUeJLs4Y4N
bmeRqxpXoMCYirG94bitywKy1lttadXLSeJxRSih698ZOG/kBDPIoUphRZFyRC+i
m0q75tieH6pDvN6T4bd+qpVrQJwXFSRT2iE3Z0X+D4roR0LlofioddCpo9H0tfrb
z4raDUkmcZTTUgBUBa16kRAQMvyBe90s2JFO6++PT6LOZ+IlfOnS+XiVvRENxAku
SOgMx/bRGNifgRhUNIH7Ed5qUBToIwk6cjLEg4Om4J0cUKYQtuFwLucLZCawtpU/
nS6ztsmyzpSaURtUUnoqQjbbLDLi7PiOQjA+V2RpcXECgYEA/rUQIf1+D5Edo2yK
n+1La5JWDr9RdM4DznbKcbFNSxYQ3K3lUnIeYz2Eqxvf9G8nhwJiqVerGJpRZeZw
BuHBnCHSQlpFq8bqzsb2cOfMexvZZToHV2L3pO/Ai/hTUeXs5Gh04WP1LbGZ2pJK
vH6MQDxo+A++Peq9TnTc7uhXnFcCgYEA1RS6G1Mnr4SOf2VL0+KK8q8rgTEpp5wX
pDnsNdY9m8njMYZRUyKVfH0Cv5x4qibNU5/RL5XkCbPOh1JTITNmhW+VjgkeUZd6
eU4w7wjtVKSiW+cbmwrsW9Rhpq26o+IT5d2y7m9W5D8j4KT9L5qgteL03SMzqMmL
WSat9MHLny8CgYEA4Bny7dUZWxz1FNrlRT0FhMomSadvRfQVSU0fZdT4Cr6Ja4Dd
KiMaNXrlBZ/q9ifugU1J/XETKvxr7dpIauWq8XKYiqTec/r6kaBhKInqUc75AaWC
3BJJjaccpIqC1KYWPgjh/YVzLRb8JWFdvGcjg0kjmk5Pti6ZDTSpRtLKctUCgYBK
b3kr/nqIl/fnjQ1WMXd0m7jI4tG4WZDwr8NSc0sGVxpkvJVAQ36RBGKnaRPF7NQh
eFztEKGeug9Vum6L1Jbl2jsWaR0MR6xjH+t8NVQjE3gcPrmoCcLTuXd7cIkYouts
i2vsWzyxc1UnLru7m3q0z1nWvmtXdUCWnip6rBBjxQKBgQC/2CuzcwLLm4u0Dtop
lG9+ztPQ8DK/2OHxbqAwyLqF9uCkaUwxPQV5tyybBsHlved5Mq6ve6zqh8Oxf13e
7SapYzbU1Ef5ihGKlAbGjB9Qt2V2jFaoSUZPYEg0XV2cMZsGxA8pjeXH0CfYM5nX
nnom6YVIODQ1vhuWQhnT+mrudg==
-----END PRIVATE KEY-----`
)


//export VerifyEncrypt1
func VerifyEncrypt1(key ,src []byte) ([]byte, error) {
	return PubEncrypt(src, key)
}

//export VerifyDecrypt1
func VerifyDecrypt1(key ,src []byte) ([]byte, error) {
	return PriDecrypt(src, key)
}

func TestOpenssl(t *testing.T)  {
	data,err:=VerifyEncrypt1([]byte(pub2),[]byte("1111"))
	fmt.Println(err)
	src,err:=VerifyDecrypt1([]byte(pri2),data)
	fmt.Println(err)
	fmt.Println(string(src))
}


// 私钥加密
func TestPriEncrypt(t *testing.T) {
	// 加密
	result, err := PriEncrypt(text, []byte(_pri1Key1024))
	if err != nil {
		panic(err)
	}
	// nRmbAgLEsFSZzieUekELhA==
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	oriData, err := PubDecrypt(result, []byte(_pub1Key1024))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(oriData))
}

// 公钥加密
func TestPubEncrypt(t *testing.T) {
	// 加密
	result, err := PubEncrypt(text, []byte(_pub2Key2048))
	if err != nil {
		panic(err)
	}
	// nRmbAgLEsFSZzieUekELhA==
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	oriData, err := PriDecrypt(result, []byte(_pri2Key2048))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(oriData))
}

