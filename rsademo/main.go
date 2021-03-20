package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"k8s.io/klog"
	"os"
)

/**
生成rsa证书的demo
*/

func main() {
	var keyLength int
	flag.IntVar(&keyLength, "l", 2048, "key length, default value 2048")
	if err := GenRsaKey(keyLength); err != nil {
		klog.Fatal("")
	}
	klog.Info("done")
}

func GenRsaKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		klog.Warning("failed to gen key")
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "privateKey",
		Bytes: derStream,
	}

	file, err := os.Create("privateKey.pem")
	if err != nil {
		klog.Warning("failed to create private key file")
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		klog.Warning("failed to write private key to file")
		return err
	}
	file.Close()

	publicKey := &privateKey.PublicKey
	derPublicStream := x509.MarshalPKCS1PublicKey(publicKey)
	block = &pem.Block{
		Type:  "publicKey",
		Bytes: derPublicStream,
	}

	file, err = os.Create("publicKey.pem")
	if err != nil {
		klog.Warning("failed to create public key file")
		return err
	}
	// 输出为pem，其实本质就是加上文集注释说明，外加将byte base64后输出到文件
	// 如果在实际应用中，rsa信息privatekey publickey可以存储到数据库，直接将byte base64后存储
	err = pem.Encode(file, block)
	if err != nil {
		klog.Warning("failed to write public key to file")
		return err
	}
	file.Close()

	return nil
}
