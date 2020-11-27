package aliyun


import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/vgmdj/utils/logger"
	"io"
	"strconv"
	"strings"
	"time"
	"utils/chars"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const  (
	endpoint = ""
	accessKeyId = ""
	accessKeySecret = ""
	bucketName = ""
)

// oss 上传文件
func PutObjectFromReader(endpoint, accessKeyId, accessKeySecret, bucketName, objectName string, objectValue io.Reader, path string) (URL string,err error) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		logger.Error(err.Error())
		return "",err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		logger.Error(err.Error())
		return "",err
	}

	timenow := time.Now().Unix()
	datenow := time.Now().Format("2006-01-02")
	URL = path + "/" + datenow + "/" + strconv.FormatInt(timenow, 10) + objectName

	// 上传文件流。
	go bucket.PutObject(URL, objectValue)

	return
}


// 上传图片 // base64
func Base64Upload(file, fileName string) (url string, err error) {
	imgAll := strings.Split(file, ";base64,")
	if len(imgAll) != 2 {
		fmt.Println(imgAll)
		return url, errors.New("文件格式不合法")
	}
	nameHead := strings.Split(imgAll[0], "/")
	if fileName == "" {
		fileName = chars.GetRandomString(20)
	}

	filename := fileName + "." + nameHead[1]
	fmt.Println(filename)
	buffer, _ := base64.StdEncoding.DecodeString(imgAll[1]) // 成图片文件并把文件写入到buffer

	pic := bytes.NewBuffer(buffer)

	// 链接 oss
	url, err = PutObjectFromReader(endpoint, accessKeyId, accessKeySecret, bucketName, filename, pic, "taofu-5g/uploads/")
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return
}