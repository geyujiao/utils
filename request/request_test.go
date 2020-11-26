package request

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
	"testing"
)

func TestPostJsonReqBytes(t *testing.T) {
	reqData := `{
		"Mobile":"17718326036",
		"Count":123,
		"Tel":"17718326040"
	}`

	resp, err := PostJsonByBytes("http://127.0.0.1:83/postjson", []byte(reqData))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("resp data", string(resp))
}

func TestPostFormData(t *testing.T) {
	resp, err := PostFormData("http://127.0.0.1:83/formdata", map[string]string{"mobile":"17718326036"})
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("resp data", string(resp))
}

func TestGet(t *testing.T) {

	url := fmt.Sprintf("http://127.0.0.1:83/get?mobile=17718326036&count=123&tel=17718326040")
	resp, err := Get(url)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("resp data", string(resp))
}
