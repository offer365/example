package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	server_crt = `
-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIJAIAevp4BX0qTMA0GCSqGSIb3DQEBCwUAMEwxCzAJBgNV
BAYTAkdCMQ4wDAYDVQQHDAVDaGluYTEUMBIGA1UECgwLZ3JwYy1zZXJ2ZXIxFzAV
BgNVBAMMDnNlcnZlci5ncnBjLmlvMB4XDTE5MTAxNTA1NTY0OFoXDTI5MTAxMjA1
NTY0OFowTDELMAkGA1UEBhMCR0IxDjAMBgNVBAcMBUNoaW5hMRQwEgYDVQQKDAtn
cnBjLXNlcnZlcjEXMBUGA1UEAwwOc2VydmVyLmdycGMuaW8wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDb6gq06bluHgJ1XT7e440SncCRfGMthuszHcjp
9vq4ywvN24lRNYEZGnMKufgDmWJ1bijWbV0/BspiZVWn5/cb9zBrGtRB3VpcXpU0
76VVPfEgxLNr9MSU947qSm+fOoymlpHScw5J74PL6dp2JLqNBh5nzWUSqQ8smiWQ
4oZuSg4NEzIZS4ilEi+RnnDZRRHpWI3GLCuLYbVHZmh8THPFjhSTlM9+OuMq/Ml5
5HWzg8Np6b8VX9ksMqkG8JkcQ4UDtFjmIylxcqmwcrzGZRpcbCUVYl28EXeQ8LB8
2aQzauJ6L+s7ZkI9ZnD2Ey8i4u7HqwU2RWvq0YPxu/Ya8ue3AgMBAAGjUDBOMB0G
A1UdDgQWBBQ+xUj1DTDqpG2sbwei9xFRDUNsKjAfBgNVHSMEGDAWgBQ+xUj1DTDq
pG2sbwei9xFRDUNsKjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBz
dJOwXWnY92I0/27VcwSWYZC2oONO8q6Mjpyg1HtAg329PmRAEXEGwqfv85KC1SBJ
jUVs/jLc4bDlWUeYtJ+xTpNbWVL7xRd7ksyY6kSxauDXCxarGjOSA2qww5JivoJW
TOK7WCXT0bOeeNdLuR+BBe7i2DtKeURC67Rn9/yl41QkboopU6pelU5Mto65ryRU
82S70uVa6GPNLk9/7F6ltM2oRAjdif0AkY25Q/sKrfg2oeqJCQLJ/SqPCAYxW24X
2dJqdaraGunnjmVfSBaSqHOZAEozGYMjUXh6w5zZEiBee3wDDTedQ46y30Vwm8le
d8cvrJQ3dIg5oc0SiBMV
-----END CERTIFICATE-----

`
	server_key = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA2+oKtOm5bh4CdV0+3uONEp3AkXxjLYbrMx3I6fb6uMsLzduJ
UTWBGRpzCrn4A5lidW4o1m1dPwbKYmVVp+f3G/cwaxrUQd1aXF6VNO+lVT3xIMSz
a/TElPeO6kpvnzqMppaR0nMOSe+Dy+nadiS6jQYeZ81lEqkPLJolkOKGbkoODRMy
GUuIpRIvkZ5w2UUR6ViNxiwri2G1R2ZofExzxY4Uk5TPfjrjKvzJeeR1s4PDaem/
FV/ZLDKpBvCZHEOFA7RY5iMpcXKpsHK8xmUaXGwlFWJdvBF3kPCwfNmkM2riei/r
O2ZCPWZw9hMvIuLux6sFNkVr6tGD8bv2GvLntwIDAQABAoIBAQDatMrLGzqP2gaF
5CM3lcRiBENULQmRaIGWx4Nds4OrrjtiEnhhLD3k0lohC8wtAClcMh8pCYDXwpAa
BT+HiflGdbJQglEf490Oyf3HtKGPwdeByD3MGpQ6tm0KctbJ23ev4UTKsCRAwZQo
gr0CDAr/X1tzzmA7i6iMqgy8J5ycJu/4qH+vUz2a5jftrcScasR+3O4pvN2W9T2i
oL+tJUPCENKPpx6ofQoepBEk5wWWK2kmHPJk8WUVwMEoNnbGWhZNwGcO8CwO+lL3
UcSDom78+KfuYInEFCfSRL/4tUvGMIVdfzz/eY/iltZ++OCs+3JZdVz/grrVNRsm
WaYrcJZZAoGBAPL2oNyXnK79kaZ9WUgtXxgrcBQc7dXhexiCejAYXT6u6FIBdpCk
D+9VEAwgIdqQpBAoudZeK8KfoCOWtPQ7sSfg+gAhAMM4h3YKLeMMrV/2stFTRFJ2
6zomFPT9tKG06WAaV8LxJ0feDEf/nMqgPhBqfcP2J9qUyoUMKwSlutLtAoGBAOe2
zr7BdFqxUXpCl2u9xKgay067X7srbGrRAMbp6CRGvZ9oVPdHMCOAFi2tT5thBdty
Bkzbv4wXQirYONn5Nw9w71A81AVP3gAxd1Gqg3V5qqZFzDxnb2YcgEcBU0ztZrpD
M2K6e9BTPHI1Xjv789zPF0n9f2s2oQaK1Oc9zZyzAoGAB1YxN1gQsCwSCOACIS7V
j0pIgSL6f5nmeK//9pHVxv3LICbRKL77iDOeX29c9lelzKMeMX34flEJqel0H2fq
CpU9l2Fnv31mgcb+6btJRPuTHMUR7BeRNNlPirJakQOAhJlnCwjzMbVf05DBcFD0
btR4ZcF6JJyXnProFaTXhmUCgYEAgslFMphA03voUEjL7O1E1dmhzYOnSh79Z+Em
PiACfo9LMnGSG6ybuD3wxsFfAIWn57AZbEJQgIMUPtiiZi3rbRTCjxh2V7U3ygYh
of/LiYAt2QHmgGWllA4cPXe7C92nsRSDKYO2pOSGZrRGxzaz83sUWxfxVpOUGfw8
pXTV3E8CgYEApABrEqYOWZjfkZJedSyhUeDYupfY9dlNpxShN3Wq+bZYLar/cYG/
fbLZwW4idWvZDRnTJNQyiD3pytTWxZLkz9H7Lt9mj18f6ZsSwNVovfGcN5aeT9sf
B1rikLP7ogmVRRQtLDfIhQZmvF1VGVInivxsoOjzkLdEwTu0twxiqxc=
-----END RSA PRIVATE KEY-----

`
)

func genServerTls(crt, key string) (tlsConfig *tls.Config) {
	cert, err := tls.X509KeyPair([]byte(crt), []byte(key))
	fmt.Println(err)
	tlsConfig = &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{cert}
	// Time returns the current time as the number of seconds since the epoch.
	// If Time is nil, TLS uses time.Now.
	tlsConfig.Time = time.Now
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.
	// The Reader must be safe for use by multiple goroutines.
	tlsConfig.Rand = rand.Reader
	return
}

func genClientTls(crt, servername string) (tlsConfig *tls.Config) {
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM([]byte(crt)) {
		return nil
	}
	return &tls.Config{ServerName: servername, RootCAs: cp}
}

type HelloServiceImpl struct{}

// 实现HelloServiceServer接口
func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *String,
) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	go client()
	var tsc credentials.TransportCredentials
	var err error
	tsc, err = credentials.NewServerTLSFromFile("server.crt", "server.key")
	fmt.Println(err)

	// 使用内嵌的证书
	_ = credentials.NewTLS(genServerTls(server_crt, server_key))
	grpcServer := grpc.NewServer(grpc.Creds(tsc))

	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}

// 这种方式，需要提前将服务器的证书告知客户端，这样客户端在链接服务器时才能进行对服务器证书认证。
// 在复杂的网络环境中，服务器证书的传输本身也是一个非常危险的问题。
// 如果在中间某个环节，服务器证书被监听或替换那么对服务器的认证也将不再可靠。
func client() {
	var tsc credentials.TransportCredentials
	var err error
	tsc, err = credentials.NewClientTLSFromFile(
		"server.crt", "server.grpc.io",
	)
	if err != nil {
		log.Fatal(err)
	}
	// 使用内嵌的证书
	credentials.NewTLS(genClientTls(server_crt, "server.grpc.io"))

	conn, err := grpc.Dial("localhost:1234",
		grpc.WithTransportCredentials(tsc),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := NewHelloServiceClient(conn)
	for i := 0; i < 10; i++ {
		reply, err := cli.Hello(context.Background(), &String{Value: "hello"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}

}
