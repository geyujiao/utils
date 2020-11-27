package sms

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
	"math/rand"
	"testing"
)

func TestSendSms(t *testing.T) {
	code := fmt.Sprintf("%v", rand.Intn(1000000))
	logger.Info("code", code)

	err :=  SendSms("LTAI4jUtvDnYmQdI", "lAUCF33krEHUjUzH4SoptjwTJWrBfx", "北京榕树科技",
		"SMS_149635151", "17718326036", fmt.Sprintf(`{"code":"%s"}`, code))
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
