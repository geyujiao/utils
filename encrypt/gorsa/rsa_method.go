package gorsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"

	"github.com/vgmdj/utils/logger"
)

// 公钥加密
func PubKeyENCTYPT(input []byte, filePath string) ([]byte, error) {
	//载入公钥
	pubKey, err := LoadPublicKet(filePath)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	pub := pubKey.(*rsa.PublicKey) //类型断言
	encryption, err := rsa.EncryptPKCS1v15(rand.Reader, pub, input)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return encryption, nil
}

// 公钥解密
func PubKeyDECRYPT(input []byte, filePath string) ([]byte, error) {
	//载入公钥
	pubKey, err := LoadPublicKet(filePath)
	if err != nil {
		logger.Error(err)
		return nil, err
	}


	decryptMsg := decrypt(pubKey.(*rsa.PublicKey), input)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return decryptMsg, nil
}

// 私钥加密
func PriKeyENCTYPT(input []byte, filePath string) ([]byte, error) {
	//载入私匙
	priKey, err := LoadPrivateKey(filePath)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	//私匙加密
	encryptMsg, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.Hash(0), input)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return encryptMsg, nil
}

// 私钥解密
func PriKeyDECRYPT(input []byte, filePath string) ([]byte, error) {
	// 载入私匙
	priKey, err := LoadPrivateKey(filePath)
	if err != nil && err.Error() == "open ./conf/private.pem: no such file or directory" {
		//logger.Error(err)
		priKey, err = LoadPrivateKey("private.pem")
		if err != nil{
			logger.Error(err)
			return nil, err
		}
	}
	// 私匙解密
	decode, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, input)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return decode, nil
}
