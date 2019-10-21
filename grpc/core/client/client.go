package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"time"
)

const (
	Server_crt = `-----BEGIN CERTIFICATE-----
MIIDAzCCAesCCQDqO1aVDNi/IzANBgkqhkiG9w0BAQsFADBDMQswCQYDVQQGEwJH
QjEOMAwGA1UEBwwFQ2hpbmExDzANBgNVBAoMBmdvYm9vazETMBEGA1UEAwwKZ2l0
aHViLmNvbTAgFw0xOTEwMjAxMTAyMzhaGA8yMTE5MDkyNjExMDIzOFowQjELMAkG
A1UEBhMCR0IxDjAMBgNVBAcMBUNoaW5hMQ8wDQYDVQQKDAZzZXJ2ZXIxEjAQBgNV
BAMMCXNlcnZlci5pbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMRF
1mgbKYNO2X0iqX89Rrzc+xftqegQ+7V0n9Sa1HE07xQVcgu05faGB4B/29HPQ/gh
JMt1IxkXlISNuQwIDM5XVSahkH1OhQmtQnTKjYXFgboRFHMQUk26lKoIZ3o9AJ8s
QTPCLBw7a9StBpeWhBzEumDymP60hmGhTft4tbY85MrmObfTZ8KbQiHIy22jXNGV
N5ok61q4tlMV8HYK89q4WX7IcQusdK9NNwL1jZNQ4+WICEe2/zs8xY9r4REONKoM
HOME5aS+EvQSVwh5LyvNuPxa8io83EjokT3yRqZllvmXD/hVS/BCM927fgsiDfm0
ezuE5+AGiMR1N0agv1cCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAHxXVd/v7noVZ
LJ8IsLty3BjMX7ZjSvkyyrchxdQQIfCoMc/UGkDZ5TgvdPkE8eAfdSVwwrcpGf8V
C4ccB9flekd6HrO7Uo9mWrKcjyQn2MjwAZNDhcs5Sxrz8TusJEQk4iQYSq0oc4Nr
qGrR/2kXEirwXi/xQ0saVXalfhkK5W+rO/YWTc8K3leARQ6BDjGbF2BHRtj6HEZL
RnhJEbx+BvplXMlWQ5CBBYt/NQa/MKJDd2stT70Si8E1lIGIGaVQAy43uT7xy8XW
jSrruOAv1SVLovhSxjsMiu/jXwZsVAtaFAuT4ajiWQHzbNqUjVnt7dJIWJPCnL6h
lhaV0MUy7Q==
-----END CERTIFICATE-----

`
	Server_key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAxEXWaBspg07ZfSKpfz1GvNz7F+2p6BD7tXSf1JrUcTTvFBVy
C7Tl9oYHgH/b0c9D+CEky3UjGReUhI25DAgMzldVJqGQfU6FCa1CdMqNhcWBuhEU
cxBSTbqUqghnej0AnyxBM8IsHDtr1K0Gl5aEHMS6YPKY/rSGYaFN+3i1tjzkyuY5
t9NnwptCIcjLbaNc0ZU3miTrWri2UxXwdgrz2rhZfshxC6x0r003AvWNk1Dj5YgI
R7b/OzzFj2vhEQ40qgwc4wTlpL4S9BJXCHkvK824/FryKjzcSOiRPfJGpmWW+ZcP
+FVL8EIz3bt+CyIN+bR7O4Tn4AaIxHU3RqC/VwIDAQABAoIBAAQbHeghoVWw4ZXf
ksIpqwAqc0pF24cSS+G45dsRvh38KIA4DqG2EBV/KksC4bta5aYcM2PaOHi+6Il5
WYSp6nKqmwpq2NX2PYw9RqWg0yMYRaV50/6wObiMja2c7WU+P3QU/ewyRK/2gkP5
tqiXKn5bkzaR/KdfaWxDbpkzJkIArLAELqEBuS0noxikrfypPanGnXk7IDhYo+rZ
cE0UHOhpkeo7gXeVc9tU/cjTRwBK7awKLIDWyknHGrL28nxMqKf06SzxG2oz6Hn3
twOtwAUS7tjophOZ6WCStgCOVFf0Ue6yJmja9xgWy/r2jJsH5/VV0xJZvmWGxr8T
IQh4oskCgYEA92Katy0Cvl1kS1/cf0ExMtOzXIwtCDu35axGl1FR3VMcoboPmH2h
HrRxSpcIgkRXz7wxsj3zttBXu8assjmwtCWzbDIE0YGYQ3v1CwDITihAyhevhW4b
UxN181RhMo1qHIcgULsVR5+P857FAHRSWWewh77ZK7x17fdQJshZujMCgYEAyxuT
R1CthfC7rbIX359tD9jb1XtG+XCgygZYv+6uoknmWMMmUqgDmQ3u8p4kuHudB6gm
/kZkxrluwJM5B8UKC1NRkejHP2ZO8ygpEGQp7t1H3BBFSfUVlu+YmfD5SjHhK9U5
2t+hfyuO8m0r+XdYk6lliEYufVlPMzJffT3rSk0CgYEAhs+jRGMw9ZBrUXAB9w8N
wob/XVW+TJhOlMiXB2r3U8cw+SktyonbvaHTgzRfHK4ltDz4UAvWvi83QEr6XX12
wBUze6ieW5Vl5pCsbryUa5MgC4Fw0yO3nEQkqN+4wBW0V6uDfrsU050ukzJYZPD+
113cI31rV5wyH+YANcJEs2UCgYEAmh0SY8qT4E4KGoJIGyadWqjyJcqk0CDl4GVw
cjJp0DrCzhdFvPI/yKMJ7I6Szmj9fhHZhJdlYGTT5MvROlQIiw9tlYlLpo+62EZg
4k8egmDlZdXyvWt6Nk0XPbfbcLDoapogjDOkFxq2HL054NDuJR0kLYMTQ4nAztgq
HJ4fKwECgYEAinsJM6lw9m3eyRRuPRFE4jNwg5KmZRjVuZ06+nPW/Sb7GXdN+5e6
62y87e63MRTm1r2C4g/esnqAOcS6iRHQtdTFrG8DU/j9F+uaB5TWZTroxqQ6h2F0
OjGZcdCMohluWRztbas01OZKSoDx1pEfP+H4kKFJ8LhWQXLU0lWibEw=
-----END RSA PRIVATE KEY-----

`
	Client_crt = `-----BEGIN CERTIFICATE-----
MIIDAzCCAesCCQDqO1aVDNi/JDANBgkqhkiG9w0BAQsFADBDMQswCQYDVQQGEwJH
QjEOMAwGA1UEBwwFQ2hpbmExDzANBgNVBAoMBmdvYm9vazETMBEGA1UEAwwKZ2l0
aHViLmNvbTAgFw0xOTEwMjAxMTAyMzhaGA8yMTE5MDkyNjExMDIzOFowQjELMAkG
A1UEBhMCR0IxDjAMBgNVBAcMBUNoaW5hMQ8wDQYDVQQKDAZjbGllbnQxEjAQBgNV
BAMMCWNsaWVudC5pbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAObU
dbPa2MMnp7X0P6TeUM9+gJgRgVdrOm05EPnf4p1xEmFq09bGupZpD+pVoU/yH/oi
wA4gwYtgk5ETtTfTbF8YUma6LYDye2m98zXiyVWpTs9pmxVRUTcpnjmyIS7mXSNE
hShN26OCTk8DtlL9STFnFWQY2Sb9PVjwDWTrXkHalQU3PFEmoQ/QPbTbBN2gydDn
WkK6LxgTaSA9xMw/j5upZh58aoLVwd8IevzKn/YnwQBEC0ptVQGl6B5EUKabhTWh
q6c4gDAhcqhdRFZa4UMcOZnzgwEuR7XzJlTL3AanBXJu5sUjDPTweOhENcdSHBQ/
sX6Cr9NFRm6bQqOrjmsCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAD4IIkxuNITUM
bHU2ebLEMPq8Udhcl9s3mBlaWf3ecDi4Yu+MBTy+ggcRhnq7zqaRVaRxdhtyVVIA
3hFwrWZK38jPGKrI9qZLoqQJu3RFq241jOjVol6zAkYhuqvO8n9AKhShjUFHPfA+
TN1BC8qb30lwYZnELaHdKFM16f7uska2lMY6N8uYqySNWFz/B77zIqUACRnvyGfS
gJ8QRDcGAjA0+SEMKtI0tB0qsL4c+de8uPaUjyO5uzWLXJap50gBi5Zx17YE8aMk
wntulWmvYwJSokZIOVi+3PDSc8Zh2ukhddA716NF2U8c7N++BZFBGcGLZK86Hh76
HayLryJQmA==
-----END CERTIFICATE-----

`
	Client_key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEA5tR1s9rYwyentfQ/pN5Qz36AmBGBV2s6bTkQ+d/inXESYWrT
1sa6lmkP6lWhT/If+iLADiDBi2CTkRO1N9NsXxhSZrotgPJ7ab3zNeLJValOz2mb
FVFRNymeObIhLuZdI0SFKE3bo4JOTwO2Uv1JMWcVZBjZJv09WPANZOteQdqVBTc8
USahD9A9tNsE3aDJ0OdaQrovGBNpID3EzD+Pm6lmHnxqgtXB3wh6/Mqf9ifBAEQL
Sm1VAaXoHkRQppuFNaGrpziAMCFyqF1EVlrhQxw5mfODAS5HtfMmVMvcBqcFcm7m
xSMM9PB46EQ1x1IcFD+xfoKv00VGbptCo6uOawIDAQABAoIBAQCTFY5qrGiy8fHL
33cudvrHPLR0MbNZINp5/oLytdaQvBwaNxgFI1yBuzCJAUdoyb/Wg44dcoHhbgiZ
yRUQHYhQkA7xpnCYWeqJ1p/DFl90Vg4B3CkVzFsT61EHMpoyaFewwViX9gSei8ma
T6M9/mdFM4pN3geA8JzGry/ZvqCxFID3Sz4/08zq9UjS54GiZgJb3lyGazdDk3Gf
h2NukbBRtvdh8iILjEM38czgqTBrDqXlFa5q0p9oq+UPn9twaVZcJ9t4IrcWIgaD
l9cYRE/agXj0cRO/IVOi/RB0e/NLiR0XqXSo4Rx7uGcSJys1yuPt96OGMIh2+c99
VGJbzBsBAoGBAP8qagFe5kJNrjweo9yNhs0H/TFx1mhCqQKPNFouDtttaCDcNvXx
d3B5KYKgWpTJPaZ1eGfPeA4OTLhKCLVG7EVQXUUsztyDS1JpuUJkm1texA30g0sw
UWhLfQfFEgWCaIQkbqZv1J5OYrc2xvPqjHfP+NG1CAte1w5QQ7FA541fAoGBAOeV
rO0yF30sDOUXlixfKN35j2FIgVB0DOT6nkpPyh1OYcdcshu3utGqmOiN7twqwyiL
m3Uucix/JbTb2m+HSAX9/s/SHHOoXeUp21wVSGYesknrBEZt3VifINzu/OFCjLk1
Plx4F0am0WrsDnAtQwgpCV29lgQjmFsXQZlUW051AoGBAJdvpbAgkUmCbsixapCn
0fv3JNZmeFgyT7n8IZbvxNOHkAgIifnXEArJbdBfuMKa2KLlDsuVfuvgormw/pAP
goP0mRZH7JFEvrwvkMqNiQJmMLcTiaRjDb13J8InvHVWmw7pzF2s+yPk44NW2CbE
6g7leAeFiDuvUrTk//e/zGzDAoGBAJx4TLaWubghIzVGkni4cuxHydB5JKYvQucT
Tg/3iR/z7ay9vLltkhRHp7i47UJkwieK7CZok0vtPJTOVvAz/z3NN3VDCWY7w/Uq
KsQ0vQ4Cf4Ph/ql3Ya6XFaUw9Dtes6YPi2r+2PsriyMrCzZP3pKM538msU1qn24s
cG4gyPBhAoGBAL+VTkIaLK07qChlT0Y2hwbmfLwAlOrPguJps7D0C0aBUDPXylOO
S5myV8jp+htbP6Mn5MEzZHhvoVSEe9GiCv9E5KMisJjPtQRRRKGNPAnTt9KJVQ1U
BCggzbZzimK/EFR72woV0071B93C4jO07jEmvkCb3gzmyWkgjREZQusj
-----END RSA PRIVATE KEY-----

`
	Ca_crt = `-----BEGIN CERTIFICATE-----
MIIDWzCCAkOgAwIBAgIJAKst9d2m1o1TMA0GCSqGSIb3DQEBCwUAMEMxCzAJBgNV
BAYTAkdCMQ4wDAYDVQQHDAVDaGluYTEPMA0GA1UECgwGZ29ib29rMRMwEQYDVQQD
DApnaXRodWIuY29tMCAXDTE5MTAyMDExMDIzOFoYDzIxMTkwOTI2MTEwMjM4WjBD
MQswCQYDVQQGEwJHQjEOMAwGA1UEBwwFQ2hpbmExDzANBgNVBAoMBmdvYm9vazET
MBEGA1UEAwwKZ2l0aHViLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBALh9e8OuCRP0zMYjbCqUk5b+d6J3tC9INL3P0VwcmWx5jCpUQLz0SGafnIL8
LworJfQkbDdOKNol9Zt4vzsxdV1k2VaZuAY0qWG5Kg+n1tCml46By9mQH3B6ngKe
cNDdBmRGYYDkuqI9g8UBgRYT4TbIQJ1Ns4wuKQR02/kCUfWypvE+8bEQEXTRKcHo
inILmcO7RvhWkfwWVbTpUv7M7K8wwIGKawDgl3DeW5g+tss0PD/iCdMo0DMRHykx
4KeTsrPYdxpxgf42LwG0aJ+/28GzYCQ4mYJaTADr5pp+vlUZWtYK8m7fFXbpGlrU
5aLTA5aEPdIuyTa2/DZXl4JBxTkCAwEAAaNQME4wHQYDVR0OBBYEFCikHb0Ms/7f
jci0C5Amwvf7cFmYMB8GA1UdIwQYMBaAFCikHb0Ms/7fjci0C5Amwvf7cFmYMAwG
A1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAG+jH3wHkTqx97/9voaftE/b
0tkbV+9b3SxPv5KoW0fm24x6UNrMPE9APt0J00Vlv20LNc/tOWquyKGDIhhe29/x
ehte/l7doGVW0Wg3xQtiIT5aJdMHNy+bSLogzV5D5sbHcPStKNj3M1wwhMj03YZ7
Nj5ua/c5aUU+MBMv0C/FNPnB+GSeRO2MxYHsZP2mBEJaLhPZ+K29kwGPCVWIESCH
IOS/jew/kfpPLavuvyPqoGAfc1xpe6QQXZUEGCtzTDU/rl/hQWMxCJg85E1S5Usx
gahmAgIzeyFCjb2txOo65VtLM0DfzzkIX2PrLz7CyiXP40m8uBMtCDG+IZS0arQ=
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

func NewRpcClient(opts ...Option) (cc *grpc.ClientConn, err error) {
	var cred credentials.TransportCredentials
	op := DefaultOpts()
	for _, opt := range opts {
		opt(op)
	}
	if op.addr == "" || op.servername == "" {
		err = errors.New("addr or servername can not be empty")
		return
	}
	// 不使用安全连接
	if !op.security {
		op.dialOption = append(op.dialOption, grpc.WithInsecure())
		return grpc.Dial(op.addr, op.dialOption...)
	}

	switch {
	case op.certFile != "" && op.keyFile != "":
		cred, err = GetClientCredByFile(op.certFile, op.keyFile, op.caFile, op.servername)
	// 使用内嵌的证书
	case op.cert != nil && len(op.cert) > 0 && op.key != nil && len(op.key) > 0:
		cred, err = GetClientCredByBytes(op.cert, op.key, op.ca, op.servername)
	default:
		err = errors.New("certificate option error")
		return
	}
	op.dialOption = append(op.dialOption, grpc.WithTransportCredentials(cred))
	return grpc.DialContext(op.ctx,op.addr, op.dialOption...)
}

type Options struct {
	servername string
	addr       string
	security   bool
	certFile   string
	keyFile    string
	caFile     string
	cert       []byte
	key        []byte
	ca         []byte
	ctx        context.Context
	dialOption []grpc.DialOption
}

type Option func(opts *Options)

func DefaultOpts() *Options {
	return &Options{
		servername: "server.io",
		security:   true,
		addr:       "127.0.0.1:1234",
		cert:       []byte(Client_crt),
		key:        []byte(Client_key),
		ca:         []byte(Ca_crt),
		ctx:        context.Background(),
	}
}

func WithServerName(servername string) Option {
	return func(opts *Options) {
		opts.servername = servername
	}
}

func WithTimeout(duration time.Duration) Option {
	return func(opts *Options) {
		opts.ctx, _ = context.WithTimeout(context.Background(), duration)
	}
}

func WithDialOption(dialOption ...grpc.DialOption) Option {
	return func(opts *Options) {
		opts.dialOption = dialOption
	}
}

func WithSecurity(security bool) Option {
	return func(opts *Options) {
		opts.security = security
	}
}

func WithAddr(addr string) Option {
	return func(opts *Options) {
		opts.addr = addr
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
