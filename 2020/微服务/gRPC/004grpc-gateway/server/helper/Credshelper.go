package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

const (
	ClientCrtFile = "keys/client.pem"
	ClientKeyFile = "keys/client.key"
	CaPem         = "keys/ca.pem"
	ServerCrtFile = "keys/server.pem"
	ServerKeyFile = "keys/server.key"

)

func GetServerCred() credentials.TransportCredentials{
	keyPair, err := tls.LoadX509KeyPair(ServerCrtFile, ServerKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	caFile, err := ioutil.ReadFile(CaPem)
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(caFile)
	newTLS := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{keyPair}, // 服务端证书
		ClientAuth:   tls.RequireAnyClientCert,  // 客户端验证需要证书，双向验证
		ClientCAs:    certPool, // 证书池
	})
	return newTLS
}

func GetClientCred() credentials.TransportCredentials {
	// ca 验证
	keyPair, err := tls.LoadX509KeyPair(ClientCrtFile, ClientKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	// 创建证书池
	certPool := x509.NewCertPool()
	caFile, _ := ioutil.ReadFile(CaPem)
	certPool.AppendCertsFromPEM(caFile)
	transportCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{keyPair}, // 客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return transportCredentials
}
