package sms

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
	"math/rand"
	"testing"
)

const  (
	accessKeyId = ""
	accessKeySecret = ""
	signName = ""
	templateCode = ""
)

func TestSendSms(t *testing.T) {
	code := fmt.Sprintf("%v", rand.Intn(1000000))
	logger.Info("code", code)

	err :=  SendSms(accessKeyId, accessKeySecret, signName,
		templateCode, "17718326036", fmt.Sprintf(`{"code":"%s"}`, code))
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
