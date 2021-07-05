package image

import (
	"bytes"
	"fmt"
	"github.com/vgmdj/utils/logger"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func PutObjectFromReader(endpoint, accessKeyId, accessKeySecret, bucketName, objectName string, objectValue io.Reader, path string) (URL string, err error) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
		return "", err
	}

	timenow := time.Now().Unix()
	datenow := time.Now().Format("2006-01-02")
	URL = path + datenow + "/" + strconv.FormatInt(timenow, 10) + objectName

	// 上传文件流。
	go bucket.PutObject(URL, objectValue)

	return
}

func PutObjectFromReaderBase64(endpoint, accessKeyId, accessKeySecret, bucketName, objectName string, objectValue io.Reader, path string) (URL string, err error) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
		return "", err
	}

	bucket, _ := client.Bucket(bucketName)
	timenow := time.Now().Unix()
	datenow := time.Now().Format("2006-01-02")
	URL = path + datenow + "/" + strconv.FormatInt(timenow, 10) + objectName

	err = bucket.PutObject(URL, objectValue) //上传读取的文件

	return
}

//删除指定object
func DeleteObject(endpoint, accessKeyId, accessKeySecret, bucketName, objectName string) (err error) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
		return err
	}
	// 删除单个文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		handleError(err)
		return err
	}
	return
}

func DownLoadObject(endpoint, accessKeyId, accessKeySecret, bucketName, objectName, downloadedFileName string) (err error) {

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 下载文件。
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		handleError(err)
	}

	return err
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

/*
压缩图片(jpeg  png)
*/
// jpeg
func CompressImageResourceJpeg(data []byte) []byte {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		logger.Error(err.Error())
		return data
	}
	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 40})
	if err != nil {
		logger.Error(err.Error())
		return data
	}
	if buf.Len() > len(data) {
		logger.Error("buf.Len() ", buf.Len(), "len(data)", len(data))
		return data
	}
	return buf.Bytes()
}

// png 不能用
func CompressImageResourcePng(data []byte) []byte {
	imgSrc, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		logger.Error(err.Error())
		return []byte{}
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 40})
	if err != nil {
		logger.Error(err.Error())
		return []byte{}
	}
	if buf.Len() > len(data) {
		logger.Error("buf.Len() ", buf.Len(), "len(data)", len(data))
		return []byte{}
	}
	return buf.Bytes()
}

func ImageUrlToBase64(imgUrl string) []byte {
	//获取远端图片
	res, err := http.Get(imgUrl)
	if err != nil {
		logger.Error("A error occurred!")
		return []byte{}
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)

	return data
	//imageBase64 := base64.StdEncoding.EncodeToString(data)
	////logger.Info("base64", imageBase64)
	//return imageBase64
}
