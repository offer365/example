package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

const (
	server_crt = `
-----BEGIN CERTIFICATE-----
MIIDATCCAekCCQDZPOavuD1IWjANBgkqhkiG9w0BAQsFADBDMQswCQYDVQQGEwJH
QjEOMAwGA1UEBwwFQ2hpbmExDzANBgNVBAoMBmdvYm9vazETMBEGA1UEAwwKZ2l0
aHViLmNvbTAeFw0xOTEwMTUwNjIzNTBaFw0yOTEwMTIwNjIzNTBaMEIxCzAJBgNV
BAYTAkdCMQ4wDAYDVQQHDAVDaGluYTEPMA0GA1UECgwGc2VydmVyMRIwEAYDVQQD
DAlzZXJ2ZXIuaW8wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC4jLpx
xwNuVAq/p0cuQkVyZQzquM6bshvEckGlNobxS/h0rkPp+efzFY156VcgJWTpdX+U
YgNxCKp2Gx0YWRV8lOeemeQ8KV/wcM/UYa8QQ8FAFzxGU+z8VKPCGKiOTkvYWFK6
QJeBXbiexvIa2KQ0h1iixRxQ9iGdQ+IeDGes9vWItlTg93DkNH66s+HbjtSFibt1
MXBCaPuELhp2b8YBu1ghWL8D4LjjpnAKm1egvxVb7hVx9UuaO0h1d5fTmLVhO5Py
yiYAQttHsVCRKBWtttbN1lkLYFDNWGTpHuqe2fyUahk25nhvHAz4kay4cDdJSdJD
Pz9R7yPitGCjt8PDAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAFhqK0BtXxUKlHZA
smWbPKWQSFLIx3MNF4+ZsYJhLRRDuLT8wHXLWiHurTeakHCjVO4Q8dzJflIIlEOE
p/EUmGGYsspWVXCifvBt+S54bzKytqoAkpHvDFxqa/XUuCnTEwXBvLoXhtFwn4eF
wmlsSq6a+Oq6auqjCDpIZHRuX+8TbwF+uJ7/VOdTf75mod1ey0+ksBeZeIllSoju
r1W7o7RCQkP8oBk9cB2yWVqGyJCyUg64TUBr7mBSmSZkyW2XZFz9jNTx8ERbHiNe
0Id28Z5oR75NE5TQxgHeJoo8wKLcT3nPNqdbDnRDjmeMPyBFZPtcT2hoSOjhq/HR
C2yiCK8=
-----END CERTIFICATE-----
`
	server_key = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAuIy6cccDblQKv6dHLkJFcmUM6rjOm7IbxHJBpTaG8Uv4dK5D
6fnn8xWNeelXICVk6XV/lGIDcQiqdhsdGFkVfJTnnpnkPClf8HDP1GGvEEPBQBc8
RlPs/FSjwhiojk5L2FhSukCXgV24nsbyGtikNIdYosUcUPYhnUPiHgxnrPb1iLZU
4Pdw5DR+urPh247UhYm7dTFwQmj7hC4adm/GAbtYIVi/A+C446ZwCptXoL8VW+4V
cfVLmjtIdXeX05i1YTuT8somAELbR7FQkSgVrbbWzdZZC2BQzVhk6R7qntn8lGoZ
NuZ4bxwM+JGsuHA3SUnSQz8/Ue8j4rRgo7fDwwIDAQABAoIBACDCEiI6EjjWQCYu
2iqy5sLcfwN3FG18mwMuyMo7uI5CTvLfL/zrOCnk7Hz0V1vEZ3otOh0rnLLGuANI
4sw8m9zGTarQZDvtmDMEw09Om5WGgVoQBcnyAcH0gDDm0ZyY90uKufyHlJ0I9slw
RO9Q/yy3zZru5AzW87aHoy50Qz7nNftNFMuj0TtjFoPaYR0gCV7lPUY+EnpCYrpc
w0qhdm5brVklYzh0QhORHA3SJBw2JLmSka8Ju/OmaYqY5WsOpbwhILhXFhhX8v3S
/q0xJryMgV2HTu6MDrkLCUCkpae/EHpIWkDVc5RYSckUB5X6vqxEpOcx2OZSLodn
qu0pvwECgYEA5dQdhul7e7LK6tFVkUBFxLGE0LlRi32RfMY5jOExoCDiEPLEAs3X
l2pvMXJGXDkRrQ9pwhdI+bH2SLbUoJjWo9GF35WiBD0yM88Otml5eWgHQgavW5Ma
K3a2jcEet1W9Y1+gStWWCagFvsQvJwhB0KHAtLMrNstnGIt615OorUMCgYEAzZCo
vCtzItRoD74g1G+C7XpOz4BcnB5d3b+/vtMWB2zqtquJ9004R1mtNPrUARjNDdCi
+Y4MfOvZOsmbrmlYFy2K3ImGnhxHInKgstePif3Ua/SUoz82NRIcFEwHBAd+sqjx
DcEdR+Wc5L35AlUAalRHUAyHm891biLpSvyL54ECgYAc6DAKjwVkCpnxLQE5Mr9T
vQw6gIScKeAJtJj6ejyWqmwku5Qh0igiuFVh/8CPyvHJNZ4Utn8MQPD8NlrKFE8y
7aCL/bMSG2xcDdgI431O4uG+0T5vIzJwcw8FB1xIrPUUMavknbawYjzOwLg5jZVR
m3a0g7CUxMKj2K9LvBvEJwKBgQCIKGYrZRg5HLnkm4nBTC9LvvSOqrYszkxcQdLu
wYBY4MLCxfJ3dJVvZS295tof8952ZRhd9cG9yLY1Iy7UIoCvsYHFu+4HsjFH1ucs
Lm2r+A4Ih5MgYhE0q88ffpAzEgfQrOgltSEA60y1kkNqUQUaaMJnejUkdhCRJ9yg
x5CfgQKBgHpjbCNNleNk76QxmFQ94lUFqF7FnqeKG9kZ+V0FtgD3Fl17OqvjfQvG
kuVbsDc8rBxFdAXz67m1CnZcE0b3Lrdw5cyb2hRvn/OjzAIQGbggV15/TaH/E21M
vxrDPZCR+ZwZawJPN8mEL5pRuEa7PNnjA07f3vS/c/X2shWH+ofU
-----END RSA PRIVATE KEY-----
`
	client_crt = `
-----BEGIN CERTIFICATE-----
MIIDATCCAekCCQDZPOavuD1IWzANBgkqhkiG9w0BAQsFADBDMQswCQYDVQQGEwJH
QjEOMAwGA1UEBwwFQ2hpbmExDzANBgNVBAoMBmdvYm9vazETMBEGA1UEAwwKZ2l0
aHViLmNvbTAeFw0xOTEwMTUwNjI0MTJaFw0yOTEwMTIwNjI0MTJaMEIxCzAJBgNV
BAYTAkdCMQ4wDAYDVQQHDAVDaGluYTEPMA0GA1UECgwGY2xpZW50MRIwEAYDVQQD
DAljbGllbnQuaW8wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCyIyAo
AsRrFfH6zBmrRBL17Vl8sO+anXWA9EBktYkc5ygs277G6Qe9DXRHiBTSH9kQ+W/3
+hM4xZM/Il2IBV0h/9Dk4j/KHAm4q1ecmBGS3FFMj778c73LtZ2HvVvol2I8UzT/
iz14Ftq4R0LIw//pp0TzPDUYsoKI92p+zZgXjyXuHqjfAInSYLWwmLVecxcVLGdJ
Iumehxo0FeMmWL5iw0NkzF5Ij49CYv6Wt4CvLHjXJO9U0xzigx3wCTejQoZ3dLQz
O91THzugRCiALRxvrDXe6zwjYfdMh/eJdqLV109vYhLETKPAp8O3wy/kXy3WLyzE
W3ouaWJ/mSMmX3fNAgMBAAEwDQYJKoZIhvcNAQELBQADggEBALc/kmOAIRc0LlBf
FM7auZQ3a8l/apSbJgWa9BLiIFuziv/gtl7C2Q0XGoNbtCM9FzZ1NkjqPgHZTbtd
IZOdLJs2O4pO37VK0O/T4rdfQ2YLY8O65l+gGYix2CSVJmmwPnD0T7wQ1SeJBWac
wStLZQz3lQLIfwmQzEi9uRh2/j7Mkb+mwk1Qn0PQvSRpmN3In2jwzzzno+eDZFOl
qZgBbvmwdHHnIK+UU+dhj8ND9gHgR/3E6CNqb0X/rDvdReXX2aMRHh7AURCAQ7Ie
8p+nqyp1uhXjPkfpWr5Vw1pw4JkvsIGVtN+jC2rpyGB7mgdrS1zhxCnLStgG5aWe
jXn0cO8=
-----END CERTIFICATE-----
`
	client_key = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAsiMgKALEaxXx+swZq0QS9e1ZfLDvmp11gPRAZLWJHOcoLNu+
xukHvQ10R4gU0h/ZEPlv9/oTOMWTPyJdiAVdIf/Q5OI/yhwJuKtXnJgRktxRTI++
/HO9y7Wdh71b6JdiPFM0/4s9eBbauEdCyMP/6adE8zw1GLKCiPdqfs2YF48l7h6o
3wCJ0mC1sJi1XnMXFSxnSSLpnocaNBXjJli+YsNDZMxeSI+PQmL+lreAryx41yTv
VNMc4oMd8Ak3o0KGd3S0MzvdUx87oEQogC0cb6w13us8I2H3TIf3iXai1ddPb2IS
xEyjwKfDt8Mv5F8t1i8sxFt6Lmlif5kjJl93zQIDAQABAoIBAQCxcmH+2TrVvUeN
X7CWNTp59dP1bL9RElbRfP2EFH2f5/fSL5drxwepX/SbqMesxILS8NaDe32YSN+z
vUTUURuD2bM5vNJ0PtfEOBIO8EBZPdRSYsKJ8bV3bdOdOpiKKfp2QyCBCi/SJ1n6
xSmWmf+bvb7mWOw/TNrRGzzfUWeW1p0Ltzy9XiwdHUc3yRaK9or8CiTWW4v6dixq
+pzLOf1WylNxN0aQtq0C/dLvCEPzvPxu8wuzyqfOPxj2Qzk6kdUqRYen5XRdWOhL
J0tP+I/bEPZxhcyw6DBKtcIUCSTU0q7lHEfhTOxMQ3uGxLqC+Au7mwk+NJ2uabGn
JKv0msOhAoGBANrRsP9YuGvSRNYQmK0/KGevdhzJpAsvqix3wodf8gVnoFdEdk3M
bwSL87WTu/ExPUMbUZrTff4Gh9iZN51LDx1yK+ZtJc+siHSmRLkfuZYGH1yJr4xt
SbVVc+mxJ2E/Wt/kGiS4aR0+rr01yAAqqnk3725Agt86jdxqMyWn4qQZAoGBANBn
1OYYib9hT4tbqC75o2McvQtLiCE/41o0Xvd5b+B2YA1Olg3xoGxex9KWFr5fQvVU
PXrCTPg04TljphjbcAw2dbU5ZedN49iAQYg7wGGKe5PA++59bmDguur3C2irJFXC
ANXLVu32BvB38wRfKaqBmqqF2LkutE5CZCGO2kfVAoGAOiZPawwgtkaClICEkkWe
by9pq+zJ808PYhHRWDhfEKChM4+2AKa7OfIXKcHAyC/Vn+e6n/JkIQWLRDwzU1GH
CsQ0dl+7FJ09BgLJcIjFwOCPpP/R7vd8BvxWeVvioy068RSk/e0jwenOdem85I5y
DxSWRC5QmRHucQyn2vHqgMECgYEAhZF2hr5FUo44n/VvliPTDsu1cY7IQZ8jxIV6
NBv1kyPrLbwnOeqZybr5UMN28i41yOxjttDe6dliXBi8tIO3jWw+Bpcx84wdMn4J
Ahphk2JhD3KJRPHJE3gU9FC/VCcT047SYDPBlCOxkN8ftraTCr+j9BRysUT4tIui
h0t6JL0CgYEAyaSvl8uZL/KTP0Y8j7xVDhSqb012tvGcGlyjrQTIjot01k4YL6a5
1ofb/jqtKYZ8klXKhrjBpVnaWVFVxQwS8AiQpMqijDO0VKvV83LOuLz1w3bpq7n4
FM0t7vikON6MPwPqWYd/gSTpkekjtDRtNoOo+yIluhvV212xlQFKqW4=
-----END RSA PRIVATE KEY-----
`
	ca_crt = `
-----BEGIN CERTIFICATE-----
MIIDWTCCAkGgAwIBAgIJAMDvNtJ7n6vRMA0GCSqGSIb3DQEBCwUAMEMxCzAJBgNV
BAYTAkdCMQ4wDAYDVQQHDAVDaGluYTEPMA0GA1UECgwGZ29ib29rMRMwEQYDVQQD
DApnaXRodWIuY29tMB4XDTE5MTAxNTA2MjMzN1oXDTI5MTAxMjA2MjMzN1owQzEL
MAkGA1UEBhMCR0IxDjAMBgNVBAcMBUNoaW5hMQ8wDQYDVQQKDAZnb2Jvb2sxEzAR
BgNVBAMMCmdpdGh1Yi5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDTe1zkUyhN1+ir0YTgUO+n7l3NPtKasZC4xZjRqxzUgC3N0YeyoKTHKT+96QPV
zqguvNbLxCP0VaF7yxm/LoYWY0yd9YEBqxyifW8Ueti5OQ63l4a9y3njyCA6Sv5f
U4m9PYJDIxnGEVfVxQTTNs4SJM2WJRu46x3QTcAb7A4Zy1y/9BfHyS3lO/AIEHq6
41yKWhAMQIgFZXgEzMTmKCVlaYWsukx1IUKpGIbZq+vuS3TnH+J3RQySgE1wpCMA
csOsRDmS64kbXXmM/Z+4wvVTZ7p+ve88deRHrx0r15ymtQoz7N1ms3C5cHA+SuJD
usHqIlTNs3qkjh7x+ut3KHVzAgMBAAGjUDBOMB0GA1UdDgQWBBQetWACAKQhB5RM
Mtu/VBEW/XhouzAfBgNVHSMEGDAWgBQetWACAKQhB5RMMtu/VBEW/XhouzAMBgNV
HRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBELfUQ/kgSXoUNnHlzuBKXAlh0
KgwZQyQf1J1B4YbCUVdlXKgui+FZhcRJFKn4uUJlj6y+dvvEy8k9o/qPnC0n0KIw
D5bsZC5V4VLl50eSer80VXfIG0hRKwiFl5/cs33jv7Sv4jkuwCGb5Z+j9ytqf4Oh
hD4noKuqCxtK3jGg+EG5H5CO8ox3/ywoVTdsZ1xc+WLJM+z+e4bxXnBgEWDkaWRd
VU7h8L6TRzBOSczDAGi1HT0W1QR6hD81We5YUKzCHl+sTI/SvOP+eK2meDaDUrs0
j+PvpuVMUPXfmeX22cNBj6sBZKcqfZhEqoYUZIbUG8Q0/+SVkSUvKhZbBBla
-----END CERTIFICATE-----
`
)

func GetClientCredByBytes(crt, key, ca []byte, servername string) (cred credentials.TransportCredentials, err error) {
	// 使用内嵌的证书
	certificate, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()

	if ca != nil && len(ca) > 0 {
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			err = errors.New("failed to append ca certs")
		}
	}
	cred = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   servername, // NOTE: 这是必需的!
		RootCAs:      certPool,
	})
	return
}

func GetClientCredByFile(crt, key, ca, servername string) (cred credentials.TransportCredentials, err error) {
	//eg: certificate, err := tls.LoadX509KeyPair("client.crt", "client.keyFile")
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		err = errors.New("failed to load cert and key file. err: " + err.Error())
	}
	certPool := x509.NewCertPool()
	if ca != "" {
		//eg: ca, err := ioutil.ReadFile("ca.crt")
		byt, err := ioutil.ReadFile(ca)
		if err != nil {
			err = errors.New("failed to read ca certs. err: " + err.Error())
		}
		if ok := certPool.AppendCertsFromPEM(byt); !ok {
			err = errors.New("failed to append ca certs")
		}
	}
	cred = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   servername, // NOTE: 这是必需的!
		RootCAs:      certPool,
	})
	return
}

func GetServerCredByBytes(crt, key, ca []byte) (cred credentials.TransportCredentials, err error) {
	// 使用内嵌的证书
	certificate, err := tls.X509KeyPair(crt, key)
	if err != nil {
		err = errors.New("failed to load cert and key file. err: " + err.Error())
	}
	certPool := x509.NewCertPool()

	if ca != nil && len(ca) > 0 {
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			err = errors.New("failed to append ca certs")
		}
	}
	cred = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: 这是可选的!
		ClientCAs:    certPool,
	})
	return
}

func GetServerCredByFile(crt, key, ca string) (cred credentials.TransportCredentials, err error) {
	//eg: certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		err = errors.New("failed to load cert and key file. err: " + err.Error())
	}
	certPool := x509.NewCertPool()
	if ca != "" {
		// eg: ca, err := ioutil.ReadFile("ca.crt")
		byt, err := ioutil.ReadFile(ca)
		if err != nil {
			err = errors.New("failed to read ca certs. err: " + err.Error())
		}
		if ok := certPool.AppendCertsFromPEM(byt); !ok {
			err = errors.New("failed to append ca certs")
		}
	}

	cred = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: 这是可选的!
		ClientCAs:    certPool,
	})
	return
}

func NewRpcServer(opts ...Option) (sv *grpc.Server, err error) {
	var cred credentials.TransportCredentials
	op := DefaultOpts()
	for _, opt := range opts {
		opt(op)
	}
	// 不使用安全连接
	if !op.security {
		sv = grpc.NewServer(op.serverOption...)
		return
	}
	switch {
	case op.certFile != "" && op.keyFile != "":
		cred, err = GetServerCredByFile(op.certFile, op.keyFile, op.caFile)
	// 使用内嵌的证书
	case op.cert != nil && len(op.cert) > 0 && op.key != nil && len(op.key) > 0:
		cred, err = GetServerCredByBytes(op.cert, op.key, op.ca)
	default:
		err = errors.New("certificate option error")
		return
	}
	op.serverOption = append(op.serverOption,grpc.Creds(cred))
	sv = grpc.NewServer(op.serverOption...)
	return
}

type Options struct {
	security bool
	certFile string
	keyFile  string
	caFile   string
	cert     []byte
	key      []byte
	ca       []byte
	serverOption []grpc.ServerOption
}

type Option func(opts *Options)

func DefaultOpts() *Options {
	return &Options{
		security: true,
		cert:     []byte(server_crt),
		key:      []byte(server_key),
		ca:       []byte(ca_crt),
	}
}

func WithSecurity(security bool) Option {
	return func(opts *Options) {
		opts.security = security
	}
}

func WithServerOption(serverOption ...grpc.ServerOption) Option {
	return func(opts *Options) {
		opts.serverOption = serverOption
	}
}

func WithCertFile(cert string) Option {
	return func(opts *Options) {
		opts.certFile = cert
	}
}

func WithKeyFile(key string) Option {
	return func(opts *Options) {
		opts.keyFile = key
	}
}

func WithCaFile(ca string) Option {
	return func(opts *Options) {
		opts.caFile = ca
	}
}

func WithCert(cert []byte) Option {
	return func(opts *Options) {
		opts.cert = cert
	}
}

func WithKey(key []byte) Option {
	return func(opts *Options) {
		opts.key = key
	}
}

func WithCa(ca []byte) Option {
	return func(opts *Options) {
		opts.ca = ca
	}
}
