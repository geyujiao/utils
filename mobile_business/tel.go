package library

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/vgmdj/utils/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
查手机号的运营商(移动-联通-电信)-归属地(河北 石家庄)
*/

type Body struct {
	Msg  string
	Code string
	Data Data
}

type Data struct {
	Mobile        string `json:"mobile"`
	NumberSection string `json:"number_section"`
	Province      string `json:"province"`
	City          string `json:"city"`
	ProvinceCode  string `json:"province_code"`
	CityCode      string `json:"city_code"`
	Supply        string `json:"supply"`
	ZipCode       string `json:"zip_code"`
}

/*
resp
{
    "msg":"success",
    "code":"200",
    "data":{
        "number_section":"1873238",
        "province":"河北",
        "city":"张家口",
        "province_code":"110",
        "city_code":"0313",
        "supply":"中国移动",
        "zip_code":"075000"
    }
}
*/

// TODO 配置文件   // 阿里云
func Tel(mobile string) (data Data, err error) {
	client := &http.Client{}
	host := "http://cowphone.market.alicloudapi.com"
	path := "/api/num/site?tel="

	resp, err := http.NewRequest("GET", host+path+mobile, nil)
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}
	resp.Header.Add("Authorization", "APPCODE f930addaadc745f3b501c3bd139c0085")

	response, err := client.Do(resp) //提交
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}
	body, err := ioutil.ReadAll(response.Body)
	logger.Info("resp body", string(body))
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}
	var TelBody = Body{}
	err = json.Unmarshal(body, &TelBody)
	if err != nil {
		logger.Info("resp body", string(body))
		logger.Error(err.Error())
		return Data{Mobile: mobile}, err
	}
	TelBody.Data.Mobile = mobile
	if TelBody.Data.Supply != "" {
		if strings.ContainsAny(TelBody.Data.Supply, "电信") {
			TelBody.Data.Supply = "电信"
		} else if strings.ContainsAny(TelBody.Data.Supply, "移动") {
			TelBody.Data.Supply = "移动"
		} else if strings.ContainsAny(TelBody.Data.Supply, "联通") {
			TelBody.Data.Supply = "联通"
		}
	}

	return TelBody.Data, err
}

func TccTel(mobile string) string {
	client := &http.Client{}
	host := "https://tcc.taobao.com/cc/json/mobile_tel_segment.htm"
	path := "?tel="

	resp, err := http.NewRequest("GET", host+path+mobile, nil)

	response, err := client.Do(resp) //提交

	fmt.Println(err)
	body, err := ioutil.ReadAll(response.Body)

	return string(body)

}

//
//__GetZoneResult_ = {
//mts:'1585078',
//province:'江苏',
//catName:'中国移动',
//telString:'15850781443',
//areaVid:'30511',
//ispVid:'3236139',
//carrier:'江苏移动'
//}
type TelTmpResp struct {
	GetZoneResult GetZoneResult `json:"__GetZoneResult_"`
}

type GetZoneResult struct {
	Province  string `json:"province"`
	CatName   string `json:"cat_name"`
	TelString string `json:"tel_string"`
}

// （淘宝） // 归属地只有省没有市
func TelTmp(mobile string) (data Data, err error) {
	logger.Info("----mobile----", mobile)
	client := &http.Client{}
	host := "http://tcc.taobao.com/cc/json/mobile_tel_segment.htm"
	path := "?tel="

	resp, err := http.NewRequest("GET", host+path+mobile, nil)
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}
	resp.Header.Add("Authorization", "APPCODE f930addaadc745f3b501c3bd139c0085")

	response, err := client.Do(resp) //提交
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err.Error())
		return data, err
	}
	b := strings.Split(string(body), "carrier:'")
	if len(b) != 2 {
		err = fmt.Errorf("body 解析失败")
		logger.Info("resp body", string(body))
		logger.Error(err.Error())
		return data, err
	}
	logger.Info(b[1])
	result := mahonia.NewDecoder("gbk").ConvertString(b[1])
	str := strings.Split(result, "'")
	fmt.Println(str[0], "======")

	logger.Info("len", len(str[0]))

	data = Data{
		Mobile: mobile,
		Supply: str[0][len(str[0])-6 : len(str[0])],
	}
	data.Province = strings.Replace(str[0], data.Supply, "", -1)
	data.Supply = "中国" + data.Supply

	return data, err
}
