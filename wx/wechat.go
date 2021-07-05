package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/vgmdj/utils/logger"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var WechatParam *GlobalWechat

const (
	//WeChatAppID     = "wx38df3d2a1f63912b"
	//WeChatAppsecret = "58f518d4e16c7ae83a64013acc2a27ab"

	WeChatAppID     = "wx2588f66f476233b2"
	WeChatAppsecret = "b50d2e7a0c0b097cdc1a6ab7d1134bd7"
)

type GlobalWechat struct {
	AccessToken   *AccessToken
	Ticket        *JsapiticketResp
	Sign          string
	SignTimestamp int64  // 时间戳
	SignNoncestr  string // 随机串
}

//签名生成规则如下：
//参与签名的字段包括 noncestr（随机字符串）,有效的jsapi_ticket, timestamp（时间戳）, url（当前网页的URL，不包含#及其后面部分） 。
//对所有待签名参数按照字段名的ASCII 码从小到大排序（字典序）后，
//使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串string1。
//这里需要注意的是所有参数名均为小写字符。对string1作sha1加密，字段名和字段值都采用原始值，不进行URL 转义。
//
func (g *GlobalWechat) GetSign(url string) {
	noncestr := RandString(16)
	timestamp := time.Now().Unix()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%v&url=%s",
		g.Ticket.Ticket, noncestr, timestamp, url)
	t := sha1.New()
	io.WriteString(t, str)
	sign := string(t.Sum(nil))

	sign = hex.EncodeToString(t.Sum(nil))
	logger.Info("sign", sign)

	g.Sign = sign
	g.SignTimestamp = timestamp
	g.SignNoncestr = noncestr
}

/*
access_token
*/
// 公众号
type AccessToken struct {
	Errcode     int
	Errmsg      string
	AccessToken string    `json:"access_token"`
	ExpiresIn   int64     `json:"expires_in"`
	Timestamp   time.Time //  // time.now
}

type JsapiticketResp struct {
	Errcode   int
	Errmsg    string
	Ticket    string
	ExpiresIn int64     `json:"expires_in"`
	Timestamp time.Time // time.now
}

func (g *GlobalWechat) GetJsapiticket() (err error) {
	if g.Ticket.Timestamp.Add(time.Duration(g.Ticket.ExpiresIn) * time.Second).
		After(time.Now().Add(5 * time.Minute)) {
		logger.Info("g.JsapiticketResp.Ticket", g.Ticket.Ticket)
	} else {

		urlstr := `https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi`
		urlstr = fmt.Sprintf(urlstr, g.AccessToken.AccessToken)
		response, err := http.Get(urlstr)
		if err == nil {
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			logger.Info("ticket resp body", string(body))

			tickeResp := new(JsapiticketResp)
			if err := json.Unmarshal(body, &tickeResp); err == nil {
				if tickeResp.Errcode != 0 {
					err = fmt.Errorf("%s", tickeResp.Errmsg)
					return err
				}
				tickeResp.Timestamp = time.Now()
				g.Ticket = tickeResp
			} else {
				logger.Error(err.Error())
				return err

			}
		} else {
			logger.Error(err.Error())
			return err

		}
	}
	return err

}

func (g *GlobalWechat) GetGlobalAccessToken() (err error) {
	if g.AccessToken.Timestamp.Add(time.Duration(g.AccessToken.ExpiresIn) * time.Second).
		After(time.Now().Add(5 * time.Minute)) {
		logger.Info("g.AccessToken", g.AccessToken.AccessToken)
	} else {
		urlstr := `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`
		urlstr = fmt.Sprintf(urlstr, WeChatAppID, WeChatAppsecret)
		response, err := http.Get(urlstr)
		if err == nil {
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				logger.Error(err.Error())
				return err

			}
			logger.Info("accesstoken resp body", string(body))

			accesstoken := new(AccessToken)
			if err := json.Unmarshal(body, &accesstoken); err == nil {
				if accesstoken.Errcode != 0 {
					err = fmt.Errorf("%s", accesstoken.Errmsg)
					return err
				}
				accesstoken.Timestamp = time.Now()
				g.AccessToken = accesstoken
			} else {
				logger.Error(err.Error())
				return err

			}
		} else {
			logger.Error(err.Error())
			return err

		}

	}
	return err

}

func init() {
	if WechatParam != nil {
		return
	}
	WechatParam = new(GlobalWechat)
	WechatParam.AccessToken = new(AccessToken)
	WechatParam.Ticket = new(JsapiticketResp)

}

//RandString 随机字符串
func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var src = rand.NewSource(time.Now().UnixNano())

	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
