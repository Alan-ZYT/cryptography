package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//创建秘钥对，自己指定位数，位数越大，越安全，但是效率越低
func generateRsaKeyPair(bit int) error {

	fmt.Println("++++++创建私钥++++++")

	//rsa
	// 1. 创建私钥, 使用GenerateKey函数产生随机数据生成器random生成一对具有指定字位数的RSA密钥。
	// example: func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error)

	priKey, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return err
	}

	//2.对私钥进行编码,生成der格式的字符串
	//x509包:公钥标准
	// example: func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
	derText, err := x509.MarshalPKCS8PrivateKey(priKey)
	if err != nil {
		return err
	}

	//3.将der格式的字符串拼装到pem格式的数据块中
	//example:
	// type Block struct {
	// 	Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
	// 	Headers map[string]string // 可选的头项
	// 	Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
	// }

	block := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil, //头信息,键值对
		Bytes:   derText,
	}

	f1, err := os.Create("rsaPriKey.pem")
	if err != nil {
		return err
	}
	defer f1.Close()

	//4. 对pem格式进行base64编码,得到最终的私钥
	// example:  err = pem.Encode(os.Stdout, &block)

	err = pem.Encode(f1, &block)
	if err != nil {
		//fmt.Println("pem Encode failed...", err)
		return err
	}

	fmt.Println("++++++创建公钥++++++")

	//1.创建公钥
	//2.通过私钥得到公钥
	pubKey := priKey.PublicKey

	// 3.对公钥进行编码，生成der格式的字符串
	//注意要使用地址，否则报错
	derText, err = x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return nil
	}

	// 4. 将der字符串拼装到pem格式的数据块中
	block = pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   derText,
	}

	// 5.对pem格式进行base64编码，得到最终的公钥
	f1, err = os.Create("rsaPublicKey.pem")
	if err != nil {
		return err
	}

	defer f1.Close()
	return pem.Encode(f1, &block)
}

func main() {
	bits := 1024
	err := generateRsaKeyPair(bits)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
