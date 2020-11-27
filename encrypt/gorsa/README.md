## gorsa

gorsa 支持rsa公钥加密私钥解密；支持rsa公钥解密私钥加密。

## gorsa 使用方法

使用方法: `go get github.com/farmerx/gorsa`

```vim
package main

import (
	"log"

	"errors"

	"github.com/farmerx/gorsa"
)
`
// 初始化设置公钥和私钥
func init() {
	if err := gorsa.RSA.SetPublicKey(Pubkey); err != nil {
		log.Fatalln(`set public key :`, err)
	}
	if err := gorsa.RSA.SetPrivateKey(Pirvatekey); err != nil {
		log.Fatalln(`set private key :`, err)
	}
}

func main() {
	// 公钥加密私钥解密
	if err := applyPubEPriD(); err != nil {
		log.Println(err)
	}
	// 公钥解密私钥加密
	if err := applyPriEPubD(); err != nil {
		log.Println(err)
	}
}

// 公钥加密私钥解密
func applyPubEPriD() error {
	pubenctypt, err := gorsa.RSA.PubKeyENCTYPT([]byte(`hello world`))
	if err != nil {
		return err
	}

	pridecrypt, err := gorsa.RSA.PriKeyDECRYPT(pubenctypt)
	if err != nil {
		return err
	}
	if string(pridecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

// 公钥解密私钥加密
func applyPriEPubD() error {
	prienctypt, err := gorsa.RSA.PriKeyENCTYPT([]byte(`hello world`))
	if err != nil {
		return err
	}

	pubdecrypt, err := gorsa.RSA.PubKeyDECRYPT(prienctypt)
	if err != nil {
		return err
	}
	if string(pubdecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

```




