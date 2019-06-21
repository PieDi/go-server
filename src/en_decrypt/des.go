package en_decrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
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
	block, err := des.NewCipher([]byte(Des.KEY))
	if err != nil {
		panic(errors.New("des key error"))
	}
	origData := []byte(orig)
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(Des.KEY))
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return string(crypted)
}

// 解密
func (Des *Des) DesDecrypt(crypted string) string {
	block, err := des.NewCipher([]byte(Des.KEY))
	if err != nil {
		panic(errors.New("des key error"))
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(Des.KEY))
	cryptedData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(cryptedData, []byte(crypted))
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
	block, err := des.NewTripleDESCipher([]byte(tripDes.KEY))
	if err != nil {
		panic(errors.New("des key error"))
	}
	origData := []byte(orig)
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(tripDes.KEY)[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return string(crypted)
}

// 3DES解密
func (tripDes *TripDes) TripleDesDecrypt(crypted string) string {
	block, err := des.NewTripleDESCipher([]byte(tripDes.KEY))
	if err != nil {
		panic(errors.New("des key error"))
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(tripDes.KEY)[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, []byte(crypted))
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData)
}














// 填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
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