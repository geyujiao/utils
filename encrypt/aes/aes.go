package library

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/vgmdj/utils/logger"
)

// AES ECB PKCS5Padding base64
func AesECBBase64Encrypt(aesKey, str []byte) (result string, err error) {
	test, err := AesECBEncrypt(str, aesKey)
	if err != nil {
		logger.Error(err.Error())
		return result, err
	}

	result = base64.StdEncoding.EncodeToString(test)

	return result, nil
}

// AES ECB PKCS5Padding base64
func AesECBBase64Decrypt(aesKey, body string) (plaint []byte, err error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		logger.Error(err)
		return
	}

	plaint, err = AesECBDecrypt(decodeBytes, []byte(aesKey))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	return
}

//加密数据  AES CBC PKCS5Padding base64
func AesCBCBase64Encrypt(key string, data []byte) (string, error) {
	encrypted, err := AesCBCEncrypt(data, []byte(key), []byte(key))
	if err != nil {
		println(err.Error())
		return "", err
	}

	result := base64.StdEncoding.EncodeToString(encrypted)
	return result, nil
}

//解密数据  AES CBC PKCS5Padding base64
func AesCBCBase64Decrypt(key string, src string) (data []byte, err error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		logger.Error(err)
		return
	}

	data, err = AesCBCDecrypt(decodeBytes, []byte(key), []byte(key))
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}

	return data, err
}

// PKCS5Padding pkcs5 padding 方式
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize //需要padding的数目
	//只要少于256就能放到一个byte中，默认的blockSize=16(即采用16*8=128, AES-128长的密钥)
	//最少填充1个byte，如果原文刚好是blocksize的整数倍，则再填充一个blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) //生成填充的文本
	return append(ciphertext, padtext...)
}

// PKCS5 UnPadding 方式
func PKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	if length == 0 {
		return originData
	}

	unpadding := int(originData[length-1])
	if unpadding > length {
		logger.Warning(fmt.Sprintf("index out of range, want to use left %d, but length is %d",
			length-unpadding, length))
		return originData
	}

	return originData[:(length - unpadding)]
}

// ZeroPadding 补0 padding
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

// ZeroUnPadding 去零 unpadding
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// AesGCMEncrypt aes gcm
func AesGCMEncrypt(plainText, key, nonce []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText = aesgcm.Seal(nil, nonce, plainText, nil)

	return

}

// AesGCMDecrypt aes gcm decrypt
func AesGCMDecrypt(cipherText, key, nonce []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err = aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return
	}

	return plainText, nil
}

// AesCBCEncrypt aes cbc pkcs7padding mode
func AesCBCEncrypt(plainText, key, iv []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding(plainText, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText = make([]byte, len(origData))

	blockMode.CryptBlocks(cipherText, origData)

	return
}

// AesCBCDecrypt aes cbc pkcs7 unpadding
func AesCBCDecrypt(cipherText, key, iv []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// CBC mode always works in whole blocks.
	if len(cipherText)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	plainText = make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	return PKCS5UnPadding(plainText), nil
}

// AesECBEncrypt aes ecb pkcs7padding mode
func AesECBEncrypt(plainText, key []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding(plainText, blockSize)

	blockMode := NewECBEncrypter(block)

	cipherText = make([]byte, len(origData))

	blockMode.CryptBlocks(cipherText, origData)

	return
}

// AesECBDecrypt aes ecb pkcs7padding mode
func AesECBDecrypt(cipherText, key []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := NewECBDecrypter(block)

	plainText = make([]byte, len(cipherText))

	blockMode.CryptBlocks(plainText, cipherText)

	return PKCS5UnPadding(plainText), nil
}

// AesCFBEncrypt aes cfb encrypt pkcs7padding mode
func AesCFBEncrypt(plainText, key, iv []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// TODO aes cfb padding 方式不一致
	//blockSize := block.BlockSize()
	//plainText =  PKCS5Padding(plainText, blockSize)

	cipherText = make([]byte, len(plainText))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText, plainText)

	return
}

// AesCFBDecrypt aes cfb encrypt pkcs7padding mode
func AesCFBDecrypt(cipherText, key, iv []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	plainText = make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	// TODO aes cfb unpadding 方式不一致
	return plainText, nil
}

// AesCTR use aes ctr
// CTR mode is the same for both encryption and decryption, but we need to solve
// the fill bytes, so split the mode into encrypt and decrypt
func AesCTREncrypt(originText, key, iv []byte) (result []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	originText = PKCS5Padding(originText, blockSize)

	result = make([]byte, len(originText))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(result, originText)

	return result, nil
}

// AesCTRDecrypt use aes ctr
// CTR mode is the same for both encryption and decryption, but we need to solve
// the fill bytes, so split the mode into encrypt and decrypt
func AesCTRDecrypt(originText, key, iv []byte) (result []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	resultText := make([]byte, len(originText))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(resultText, originText)

	return PKCS5UnPadding(resultText), nil
}

// AesOFB use aes ofb
// OFB mode is the same for both encryption and decryption, but we need to solve
// the fill bytes, so split the mode into encrypt and decrypt
func AesOFBEncrypt(originText, key, iv []byte) (result []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// TODO aes ofb padding 方式不一致
	//blockSize := block.BlockSize()
	//originText = PKCS5Padding(originText, blockSize)

	result = make([]byte, len(originText))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(result, originText)

	return
}

// AesOFB use aes ofb
// CTR mode is the same for both encryption and decryption, but we need to solve
// the fill bytes, so split the mode into encrypt and decrypt
func AesOFBDecrypt(originText, key, iv []byte) (result []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	resultText := make([]byte, len(originText))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(resultText, originText)

	// TODO aes ofb unpadding 方式不一致
	return resultText, nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
