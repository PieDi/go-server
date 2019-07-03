package en_decrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

type Des struct {
	KEY string
}

var DES *Des

func InstanceDes(key string) *Des  {
	DES := &Des{KEY:key}
	return DES
}

// DES加密
func (Des *Des) DesEncrypt(orig string) string {
	key := []byte(Des.KEY)
	iv := []byte(Des.KEY)
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	origData := []byte(orig)
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

// 解密
func (Des *Des) DesDecrypt(crypted string) string {
	key := []byte(Des.KEY)
	iv := []byte(Des.KEY)
	block, err := des.NewCipher(key)
	if err != nil {
		panic(errors.New("des key error"))
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	crypteByte, _ := base64.StdEncoding.DecodeString(crypted)
	cryptedData := make([]byte, len(crypteByte))
	// origData := crypted
	blockMode.CryptBlocks(cryptedData, crypteByte)
	cryptedData = PKCS5UnPadding(cryptedData)
	// origData = ZeroUnPadding(origData)
	return string(cryptedData)
}



type TripDes struct {
	KEY string
}

var TRIPDES *TripDes

func InstanceTripDes(key string) *TripDes  {
	TRIPDES := &TripDes{KEY:key}
	return TRIPDES
}

// 3DES加密
func (tripDes *TripDes) TripleDesEncrypt(orig string) string {
	key := []byte(tripDes.KEY)
	iv := []byte(tripDes.KEY)

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	origData := []byte(orig)
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv[:block.BlockSize()])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

// 3DES解密
func (tripDes *TripDes) TripleDesDecrypt(crypted string) string {
	key := []byte(tripDes.KEY)
	iv := []byte(tripDes.KEY)

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv[:block.BlockSize()])
	crytedByte, _ := base64.StdEncoding.DecodeString(crypted)
	//crytedByte := []byte(crypted)
	origData := make([]byte, len(crytedByte))
	// origData := crypted
	blockMode.CryptBlocks(origData, crytedByte)
	origData = PKCS5UnPadding(origData)

	// origData = ZeroUnPadding(origData)
	return string(origData)
}


// PKCS7 填充
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {

	if blocksize <= 0 {
		panic(errors.New("invalid blocksize"))
	}
	if ciphertext == nil || len(ciphertext) == 0 {
		panic(errors.New("invalid PKCS7 data (empty or not padded)"))
	}
	n := blocksize - (len(ciphertext) % blocksize)
	fmt.Println(n, len(ciphertext), blocksize)
	pb := make([]byte, len(ciphertext)+n)
	copy(pb, ciphertext)
	copy(pb[len(ciphertext):], bytes.Repeat([]byte{byte(n)}, n))
	return pb


	//padding := blocksize - len(ciphertext)%blocksize
	//padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {

	if origData == nil || len(origData) == 0 {
		panic(errors.New("invalid PKCS7 data (empty or not padded)"))
	}

	//if len(origData)%block.BlockSize() != 0 {
	//	panic(errors.New("invalid padding on input"))
	//}

	c := origData[len(origData)-1]
	n := int(c)
	if n == 0 || n > len(origData) {
		panic(errors.New("invalid padding on input"))
	}
	for i := 0; i < n; i++ {
		if origData[len(origData)-n+i] != c {
			panic(errors.New("invalid padding on input"))
		}
	}
	return origData[:len(origData)-n]
}

// PKCS5 填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	fmt.Println(padding, len(ciphertext), blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 0 填充
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}