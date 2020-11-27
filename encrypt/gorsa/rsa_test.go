package gorsa

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/vgmdj/utils/logger"
)

// 公钥加密私钥解密
func Test_PubENCTYPTPriDECRYPT(t *testing.T) {

	pubenctypt, err := PubKeyENCTYPT([]byte(`e2vpcdkmq8fyxs3b`), "public.pem")
	if err != nil {
		t.Error(err)
	}
	encodeString := base64.StdEncoding.EncodeToString(pubenctypt)
	logger.Info("public after", encodeString)

	//encodeString = "QoRbYlWF5TDafVU0VI2RuZskclw1+QHeioFMAKYRud8PbRdTG7ZBTMfkEItsfb/ufw45sqKqYmCpUjxEGsEXnWA/OaaOUTFeExmcOhQPSNHYAQsBNFj7prW5+KvqfdVCTGXV4xEve8LNJXRrfCWy+eeyiSXjxNLZVUrnfHuEQY0="
	fmt.Println("=================解码====================")
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	pridecrypt, err := PriKeyDECRYPT(decodeBytes, "private.pem")
	logger.Info(string(pridecrypt))
	if err != nil {
		t.Error(err)
	}
	if string(pridecrypt) != `e2vpcdkmq8fyxs3b` {
		t.Error(`不符合预期`)
	}
}

// 公钥解密私钥加密
func Test_PriENCTYPTPubDECRYPT(t *testing.T) {
	t.SkipNow()

	//prienctypt, err := PriKeyENCTYPT([]byte(`e2vpcdkmq8fyxs3b`))
	//if err != nil {
	//	t.Error(err)
	//}
	//pubdecrypt, err := PubKeyDECRYPT(prienctypt)
	//if err != nil {
	//	t.Error(err)
	//}
	//if string(pubdecrypt) != `e2vpcdkmq8fyxs3b` {
	//	t.Error(`不符合预期`)
	//}
}
