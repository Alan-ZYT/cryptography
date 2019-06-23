package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/*
aes + crt

1.aes
	密钥:16
	分组长度:16

2.分组模式:ctr
	不需要填充
	需要提供数字
*/

//输入明文,输出密文
func aesCtrEncrypt(plainText, key []byte) ([]byte, error) {
	//TODO
	/*
		第一步：创建aes密码接口
			aes包，golang内置的标准库
			创建一个cipher.Block接口。
			func NewCipher(key []byte) (cipher.Block, error)
			1. 参数：秘钥
			2. 返回一个分组接口
	*/
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//打印aes的分组长度
	fmt.Println("block.BlockSize: ", block.BlockSize())
	/*
		//第二步：创建分组模式ctr
		//crypto/cipher包
		// func NewCTR(block Block, iv []byte) Stream
		//参数1：填写分组接口
		//参数2：初始向量

		// iv 要与算法长度一致，16字节
		// 使用bytes.Repeat创建一个切片，长度为blockSize()，16个字符"1"
	*/
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	/*
		//第三步：加密
		// XORKeyStream(dst, src []byte)
		//参数1：密文空间
		//参数2：明文

	*/
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	// return []byte("Hello world"), nil
	return dst, nil

}

//输入密文,得到明文

func aesCtrDecrypt(encryptData, key []byte) ([]byte, error) {
	return aesCtrEncrypt(encryptData, key)
}

func main() {
	//明文，需要加密的数据
	src := "你好!"
	//src := "Stream接口代表一个流模式的加/解密器"

	//对称加密,aes 16字节
	key := "1234567887654321"
	// key := "12345678876543210" //17, 无效的
	encryptData, err := aesCtrEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("加密错误 error", err)
		return
	}

	fmt.Printf("encryptData: %x\n", encryptData)

	//调用解密函数
	plainText, err := aesCtrDecrypt(encryptData, []byte(key))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("解密后的数据: %s\n", plainText)
}
