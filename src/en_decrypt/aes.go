package en_decrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

/*
* CBC
*/
type AesCBC struct {
	KEY string
}

var AESCBC *AesCBC

func InstanceAesCBC(key string) *AesCBC {
	AESCBC := &AesCBC{key}
	return AESCBC
}

// 加密
func (aesCBC *AesCBC) CBCAesEncrypt(orig string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(aesCBC.KEY)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

// 解密
func (aesCBC *AesCBC) CBCAesDecrypt(cryted string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(aesCBC.KEY)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}
//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


/*
* CFB
*/
type AesCFB struct {
	KEY string
}

var AESCFB *AesCFB
func InstanceAesCFB(key string) *AesCFB {
	AESCFB := &AesCFB{key}
	return AESCFB
}

// 加密
func (aesCFB *AesCFB) CFBAesEncrypt(orig string) string {
	key, _ := hex.DecodeString(aesCFB.KEY)
	plaintext := []byte(orig)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x", ciphertext)
}

// 解密
func (aesCFB *AesCFB) CFBAesDecrypter(cryted string) string {
	key, _ := hex.DecodeString(aesCFB.KEY)
	ciphertext, _ := hex.DecodeString(cryted)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}