package wx

import (
	"github.com/vgmdj/utils/logger"
	"testing"
	"time"
)

func TestGetGlobalAccessToken(t *testing.T) {

	err := WechatParam.GetGlobalAccessToken()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = WechatParam.GetJsapiticket()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	WechatParam.GetSign("")
	logger.Info("accessToken", WechatParam.AccessToken)
	logger.Info("ticket", WechatParam.Ticket)
	logger.Info("WechatParam---", WechatParam)
}

func TestRandString(t *testing.T) {
	for i := 1; i < 3; i++ {
		logger.Info(i, " ---", RandString(16))
		time.Sleep(1 * time.Second)
	}
}
