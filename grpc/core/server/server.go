package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	// eg: certificate, err := tls.LoadX509KeyPair("client.crt", "client.keyFile")
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
		ServerName:   servername, // NOTE: 这是必需的!
		RootCAs:      certPool,
	})
	return
}

func GetServerCredByBytes(crt, key, ca []byte) (cred credentials.TransportCredentials, err error) {
	// 加载证书和密钥 （同时能验证证书与私钥是否匹配）
	certificate, err := tls.X509KeyPair(crt, key)
	if err != nil {
		err = errors.New("failed to load cert and key file. err: " + err.Error())
	}
	// 将根证书加入证书池
	// 测试证书的根如果不加入可信池，那么测试证书将视为不可信，无法通过验证。
	certPool := x509.NewCertPool()
	if len(ca) > 0 {
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
	// 加载证书和密钥 （同时能验证证书与私钥是否匹配）
	// eg: certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		err = errors.New("failed to load cert and key file. err: " + err.Error())
	}
	// credentials.NewServerTLSFromCert(&certificate)
	// 将根证书加入证书池
	// 测试证书的根如果不加入可信池，那么测试证书将视为不可惜，无法通过验证。
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
	op.serverOption = append(op.serverOption, grpc.Creds(cred))
	sv = grpc.NewServer(op.serverOption...)
	return
}

type Options struct {
	security     bool
	certFile     string
	keyFile      string
	caFile       string
	cert         []byte
	key          []byte
	ca           []byte
	serverOption []grpc.ServerOption
}

type Option func(opts *Options)

func DefaultOpts() *Options {
	return &Options{
		security: true,
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
