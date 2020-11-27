package sms


import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/vgmdj/utils/logger"
)

func SendSms(accessKeyId, accessSecret, signName, templateCode, phoneNumbers, templateParam string) error {
	client, err := sdk.NewClientWithAccessKey("default", accessKeyId, accessSecret)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["TemplateParam"] = templateParam
	request.QueryParams["SignName"] = signName
	request.QueryParams["TemplateCode"] = templateCode
	request.QueryParams["PhoneNumbers"] = phoneNumbers

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	result := smsResponse{}
	json.Unmarshal(response.GetHttpContentBytes(), &result)
	if result.Code != "OK" {
		logger.Info("response", response.GetHttpContentString())
		logger.Error(result.Message)
		return fmt.Errorf("%s", result.Message)
	}
	return nil
}

type smsResponse struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}

