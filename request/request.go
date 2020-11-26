package request

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/vgmdj/utils/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// post application/json
func PostJsonByBytes(url string, body []byte) (resp []byte, err error) {

	resp, err = sendRequest(url, bytes.NewReader(body), map[string]string{"Content-Type": "application/json"}, "POST")

	return
}

// post form-data
func PostFormData(url string, bodyData map[string]string) (resp []byte, err error) {
    body := new(bytes.Buffer)
    w := multipart.NewWriter(body)
    for k, v := range bodyData{
    	w.WriteField(k, v)
	}
	w.Close()

	resp, err = sendRequest(url, body, map[string]string{"Content-Type": w.FormDataContentType()}, "POST")
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}
	return resp, nil
}

// get x-www-form-urlencoded
func Get(url string) (resp []byte, err error) {
	resp, err = sendRequest(url, nil, map[string]string{"Content-Type": "x-www-form-urlencoded"}, "GET")

	return
}

// post x-www-form-urlencoded   不要用这种格式

// post form-data(file类型)
func PostFormDataFile(url string, file *multipart.FileHeader) (resp []byte, err error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile("video", file.Filename)
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}
	fd, err := file.Open()
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}
	defer fd.Close()
	_, err = io.Copy(fw, fd)
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}
	w.Close()

	resp, err = sendRequest(url, buf, map[string]string{"Content-Type": w.FormDataContentType()}, "POST")
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}

	return resp, err
}

func sendRequest(url string, body io.Reader, addHeaders map[string]string, method string) (resp []byte, err error) {
	// 创建req
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}
	// 设置headers
	for k, v := range addHeaders {
		logger.Info(k, v)
		req.Header.Add(k, v)
	}
	if addHeaders["Content-Type"] == "" {
		logger.Info("----")
		req.Header.Add("Content-Type", "application/json")
	}
	// 发送请求
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Second * 150}
	response, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("sendRequest failed, url=%s, response status code=%v", url, response.StatusCode)
		return resp, err
	}
	// 解析结果
	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err.Error())
		return resp, err
	}

	return resp, nil
}