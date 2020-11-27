package gorsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/subtle"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/vgmdj/utils/logger"
	"io/ioutil"
	"math/big"
	"os"
)

// 生成公钥 && 私钥文件
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem") // 生成到运行目录下
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem") // 生成到运行目录下
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// LoadPrivateKey载入私匙
func LoadPrivateKey(filePath string) (priKey *rsa.PrivateKey, err error) {
	// 以Pkcs8的方式载入私钥,返回私钥
	privateKeyFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	privateKey, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	blockPri, _ := pem.Decode(privateKey)
	if blockPri == nil {
		err = fmt.Errorf("私钥格式错误1:%s", privateKey)
		return
	}
	prkI, err := x509.ParsePKCS1PrivateKey([]byte(blockPri.Bytes))
	if err != nil {
		fmt.Println("err2")
		return nil, err
	}

	return prkI, err
}

// LoadPublicKet 载入公钥
func LoadPublicKet(filePath string) (priKey interface{}, err error) {
	privateKeyFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	publicKey, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		return nil, err
	}

	blockPub, _ := pem.Decode(publicKey)
	if blockPub == nil {
		return nil, errors.New("blockPub空")
	}

	pubKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		fmt.Println("err3")
		return
	}

	return pubKey, nil

}
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func decrypt(pub *rsa.PublicKey, sig []byte) []byte {
	k := (pub.N.BitLen() + 7) / 8
	m := new(big.Int)
	c := new(big.Int).SetBytes(sig)
	e := big.NewInt(int64(pub.E))
	m.Exp(c, e, pub.N)
	em := leftPad(m.Bytes(), k)
	firstByteIsZero := subtle.ConstantTimeByteEq(em[0], 0)
	secondByteIsTwo := subtle.ConstantTimeByteEq(em[1], 1)
	lookingForIndex := 1
	index := 0
	for i := 2; i < len(em); i++ {
		equals0 := subtle.ConstantTimeByteEq(em[i], 0)
		index = subtle.ConstantTimeSelect(lookingForIndex&equals0, i, index)
		lookingForIndex = subtle.ConstantTimeSelect(equals0, 0, lookingForIndex)
	}
	validPS := subtle.ConstantTimeLessOrEq(2+8, index)
	valid := firstByteIsZero & secondByteIsTwo & (^lookingForIndex & 1) & validPS
	index = subtle.ConstantTimeSelect(valid, index+1, 0)
	return em[index:]
}

func CreateCdkey(newPrivateKey bool) {
	// 载入私匙
	priKey, err := LoadPrivateKey("private.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("私匙为", priKey)

	// 载入公钥
	pubKey, err := LoadPublicKet("public.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("公匙为", pubKey)

	fmt.Println("-----------------公匙加密，私匙解密-----------------")
	// 公匙加密
	byteArr := []byte("hello world")
	pub := pubKey.(*rsa.PublicKey) // 类型断言
	encryption, err := rsa.EncryptPKCS1v15(rand.Reader, pub, byteArr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("加密结果为:", base64.StdEncoding.EncodeToString(encryption))

	// 私匙解密

	decode, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, encryption)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("解密结果为：", string(decode))
	fmt.Println("-----------------私匙加密，公匙解密-----------------")
	// 私匙加密
	encryptMsg, _ := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.Hash(0), []byte(byteArr))
	fmt.Println("加密结果为:", base64.StdEncoding.EncodeToString(encryptMsg))

	// 公钥解密
	decryptMsg := decrypt(pubKey.(*rsa.PublicKey), encryptMsg)
	fmt.Println("解密结果为:", string(decryptMsg))

}
