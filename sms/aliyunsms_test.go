package sms

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	accessKeyId     = ""
	accessKeySecret = ""
	signName        = ""
	templateCode    = ""
)

func TestSendSms(t *testing.T) {
	code := fmt.Sprintf("%v", rand.Intn(1000000))
	fmt.Println("code", code)

	err := SendSms(accessKeyId, accessKeySecret, signName,
		templateCode, "17718326036", fmt.Sprintf(`{"code":"%s"}`, code))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
