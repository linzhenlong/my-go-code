package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

const (
	CrtFile = "keys/client.pem"
	KeyFile = "keys/client.key"
	CaPem = "keys/ca.pem"
)

func GetClientCred() credentials.TransportCredentials {
	// ca 验证
	keyPair, err := tls.LoadX509KeyPair(CrtFile, KeyFile)
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

