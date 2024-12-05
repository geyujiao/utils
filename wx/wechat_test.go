package wx

import (
	"fmt"
	"testing"
	"time"
)

func TestGetGlobalAccessToken(t *testing.T) {

	err := WechatParam.GetGlobalAccessToken()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = WechatParam.GetJsapiticket()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	WechatParam.GetSign("")
	fmt.Println("accessToken", WechatParam.AccessToken)
	fmt.Println("ticket", WechatParam.Ticket)
	fmt.Println("WechatParam---", WechatParam)
}

func TestRandString(t *testing.T) {
	for i := 1; i < 3; i++ {
		fmt.Println(i, " ---", RandString(16))
		time.Sleep(1 * time.Second)
	}
}
