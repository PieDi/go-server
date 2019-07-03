package en_decrypt

import (
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
	key := []byte(aesCBC.KEY)
	iv := []byte(aesCBC.KEY)
	// 分组秘钥
	block, _ := aes.NewCipher(key)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS5Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
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
	key := []byte(aesCBC.KEY)
	iv := []byte(aesCBC.KEY)
	// 分组秘钥
	block, _ := aes.NewCipher(key)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS5UnPadding(orig)
	return string(orig)
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
	plaintext := []byte(orig)
	block, err := aes.NewCipher([]byte(aesCFB.KEY))
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
	return hex.EncodeToString(ciphertext)
}

// 解密
func (aesCFB *AesCFB) CFBAesDecrypter(cryted string) string {
	ciphertext, _ := hex.DecodeString(cryted)
	block, err := aes.NewCipher([]byte(aesCFB.KEY))
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


/*
* ECB
*/
type AesECB struct {
	KEY string
}

var AESECB *AesECB
func InstanceAesECB(key string) *AesECB {
	AESECB := &AesECB{key}
	return AESECB
}

func (aesECB *AesECB) ECBAesEncrypt(orig string) string {

	origData := []byte(orig)
	// 分组秘钥
	block, err := aes.NewCipher([]byte(aesECB.KEY))
	if err != nil{
		panic(err)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(origData) > 0 {
		// 每次运算一个block
		block.Encrypt(text, origData)
		origData = origData[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// 解密
func (aesECB *AesECB) ECBAesDecrypter(cryted string) string {

	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	// 分组秘钥
	block, _ := aes.NewCipher([]byte(aesECB.KEY))
	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	for len(crytedByte) > 0 {
		block.Decrypt(text, crytedByte)
		crytedByte = crytedByte[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}
	return string(PKCS7UnPadding(plaintext))

}
