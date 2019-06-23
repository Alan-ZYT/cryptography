package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

/*
//背景：des + cbc
//des: 秘钥：8字节，分组长度：8字节
//cbc: 1.长度与算法相同(8字节) 2. 需要填充

// En //肯定
// De, Un //否定
*/

//输入明文，输出密文
func desCBCEncrypt(plainText /*明文*/, key []byte) ([]byte, error) {
	//第一步：创建des密码接口, 输入秘钥，返回接口
	// func NewCipher(key []byte) (cipher.Block, error)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 第二步：创建cbc分组
	// 返回一个密码分组链接模式的、底层用b解密的BlockMode接口
	// func NewCBCEncrypter(b Block, iv []byte) BlockMode
	blockSize := block.BlockSize()

	// 创建一个8字节的初始化向量
	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	BlockMode := cipher.NewCBCEncrypter(block, iv)

	//第三步：填充
	//TODO
	plainText, err = paddingNumber(plainText, blockSize)
	if err != nil {
		return nil, nil
	}

	//第四步：加密
	// type BlockMode interface {
	// 	// 返回加密字节块的大小
	// 	BlockSize() int
	// 	// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
	// 	CryptBlocks(dst, src []byte)
	// }

	//密文与明文共享空间，没有额外分配
	BlockMode.CryptBlocks(plainText /*密文*/, plainText /*明文*/)

	return plainText, nil
}

//输入密文，得到明文
func desCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	//TODO
	//第一步：创建des密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, nil
	}

	//第二步：创建cbc分组
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	BlockMode := cipher.NewCBCDecrypter(block, iv)

	//第三步：解密
	BlockMode.CryptBlocks(encryptData /*明文*/, encryptData /*密文*/)

	//第四步: 去除填充
	//TODO
	encryptData, err = unPaddingNumber(encryptData)
	if err != nil {
		return nil, nil
	}

	//return []byte("Hello world"), nil
	return encryptData, nil
}

//填充数据
func paddingNumber(src []byte, blockSize int) ([]byte, error) {

	if src == nil {
		return nil, errors.New("src长度不能为空...")
	}

	fmt.Println("调用paddingNumber函数...")

	//1. 得到分组之后剩余的长度 5
	leftNumber := len(src) % blockSize //5

	//2. 得到需要的个数 8-5=3
	needNumber := blockSize - leftNumber //3

	//3. 创建一个切片,包含三个三
	newSlice := bytes.Repeat([]byte{byte((needNumber))}, needNumber) //newSlice ==>[]byte{3,3,3l}
	fmt.Printf("newSlice: %v\n", newSlice)

	//4. 将新切片追加到src
	src = append(src, newSlice...)

	return src, nil
}

//解密后去除填充数据
func unPaddingNumber(src []byte) ([]byte, error) {

	fmt.Println("调用unPaddingNumber函数...")

	//1. 获取最后一个字符
	lastChar := src[len(src)-1] //byte(3)

	//2. 将字符转换为数字
	num := int(lastChar) //int3

	//3. 截取切片(左闭右开)
	return src[:len(src)-num], nil
}

func main() {
	//src := "this is a test routine" //明文
	src := "Base64是一种基于64个可打印字符来表示二进制数据的表示方法。" //明文
	key := "12345678"                          //密钥

	//加密处理 encryptData
	encryptData, err := desCBCEncrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("encryptData: %x\n", encryptData)

	key = "12345678" //密钥
	//调用解密的函数
	plainText, err := desCBCDecrypt(encryptData, []byte(key))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("解密后的数据: %s\n", plainText)
	//fmt.Printf("解密后的数据 hex: %x\n", plainText)
}
