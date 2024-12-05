package image

import (
	"fmt"
	"testing"
)

const (
	BaseUrl        = "https://fudun.oss-cn-beijing.aliyuncs.com/"
	TestImageUrl   = "https://fudun.oss-cn-beijing.aliyuncs.com/fudun_manager/uploads/2021-03-16/1615883918rTY609K0mkx0K9taz5Vb.jpeg"
	TestImageUrl1M = "https://fudun.oss-cn-beijing.aliyuncs.com/fudun_manager/uploads/2021-04-15/1618468355tsz8iezVqLRQsOaw5DVv.png"
	TestImagePngKB = "https://fudun.oss-cn-beijing.aliyuncs.com/fudun_manager/uploads/2021-04-15/1618456878blioRbCGie53xoBoejYS.png"
)

func TestDownLoadObject(t *testing.T) {
	//config.InitConfig("")
	//err := DownLoadObject(config.OssConfig.Endpoint, config.OssConfig.AccessKeyId,
	//	config.OssConfig.AccessKeySecret, config.OssConfig.BucketName,
	//	TestImageUrl,
	//	"dilireba")
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("success")
	//}
}
func TestImageUrlToBase64(t *testing.T) {
	ImageUrlToBase64(TestImagePngKB)
}

func TestCompressImageResourceJpeg(t *testing.T) {
	base64Image := ImageUrlToBase64(TestImagePngKB)
	//image2 := CompressImageResourceJpeg([]byte(base64Image))

	//if strings.Contains(base64Image, "{"){
	//	fmt.Println("have")
	//	strs := strings.Split(base64Image, "{")
	//	if len(strs) == 2{
	//		base64Image = strs[0]
	//	}
	//}
	fmt.Println("base64 2---", base64Image)
	image2 := CompressImageResourcePng([]byte(base64Image))

	fmt.Println("image2", string(image2))
}

func TestBase64Upload(t *testing.T) {
	////showTips()
	//execute()
	//time.Sleep(5 * time.Minute) /** 如果不是自己点击退出，延时5分钟 */
	base64Image := ImageUrlToBase64(TestImagePngKB)

	Png(base64Image, "headpng2.png")
}
