package en_decrypt

import (
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	_ "encoding/base64"
	"encoding/pem"
	"errors"
)

type Rsa struct {
	PRIKEY, PUBKEY string   // 私钥
}

var RSA *Rsa

func InstanceRsa(priKey, pubKey string) *Rsa {
	RSA := &Rsa{priKey, pubKey}
	return RSA
}

// 加密
func (Rsa *Rsa)RsaEnceypt(orig string) string {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(Rsa.PUBKEY))
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	encByte, encErr := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(orig))
	if encErr != nil {
		panic(encErr)
	}
	return base64.StdEncoding.EncodeToString(encByte)
}

// 解密
func (Rsa *Rsa) RsaDecrypt(crypt string) string {
	//解密
	block, _ := pem.Decode([]byte(Rsa.PRIKEY))
	if block == nil {
		panic(errors.New("private key error"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	cryptByte, _ := base64.StdEncoding.DecodeString(crypt)
	// 解密
	decByte, decErr := rsa.DecryptPKCS1v15(rand.Reader, priv, cryptByte)
	if decErr != nil {
		panic(decErr)
	}
	return string(decByte)
}