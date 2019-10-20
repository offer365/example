## 生成服务器和客户端分别生成私钥和证书
> 用以下命令为服务器和客户端分别生成私钥和证书
```bash
openssl genrsa -out server.key 2048
openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
    -key server.key -out server.crt

openssl genrsa -out client.key 2048
openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
    -key client.key -out client.crt
```

   以上命令将生成server.key、server.crt、client.key和client.crt四个文件。
   
   其中以.key为后缀名的是私钥文件，需要妥善保管。
   
   以.crt为后缀名是证书文件，也可以简单理解为公钥文件，并不需要秘密保存。
   
   在subj参数中的/CN=server.grpc.io表示服务器的名字为server.grpc.io，在验证服务器的证书时需要用到该信息。

>  **以上这种方式，需要提前将服务器的证书告知客户端，这样客户端在链接服务器时才能进行对服务器证书认证。在复杂的网络环境中，服务器证书的传输本身也是一个非常危险的问题。如果在中间某个环节，服务器证书被监听或替换那么对服务器的认证也将不再可靠。**


> 为了避免证书的传递过程中被篡改，可以通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名。这样客户端或服务器在收到对方的证书后可以通过根证书进行验证证书的有效性。。

## 生成根证书
>根证书的生成方式和自签名证书的生成方式类似:
```bash
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=gobook/CN=github.com" \
    -key ca.key -out ca.crt
```

## 重新对服务器端证书进行签名

```bash
openssl req -new \
    -subj "/C=GB/L=China/O=server/CN=server.io" \
    -key server.key \
    -out server.csr
openssl x509 -req -sha256 \
    -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
    -in server.csr \
    -out server.crt
```

## 重新对客户端证书进行签名

> 如果客户端的证书也采用CA根证书签名的话，服务器端也可以对客户端进行证书认证。我们用CA根证书对客户端证书签名：
```bash
openssl req -new \
    -subj "/C=GB/L=China/O=client/CN=client.io" \
    -key client.key \
    -out client.csr
openssl x509 -req -sha256 \
    -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
    -in client.csr \
    -out client.crt
```