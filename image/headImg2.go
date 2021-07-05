package image

import (
	"bytes"
	//"encoding/json"
	"fmt"
	//"os"
	//"io/ioutil"
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
)

func Png(file_origin []byte, toFilePath string) bool {
	origin, err := png.Decode(bytes.NewBuffer(file_origin))
	if err != nil {
		fmt.Println("png.Decode(file_origin)")
		log.Fatal(err)
		return false
	}
	//temp, err := os.Open(originFilePath)
	//if err != nil {
	//	fmt.Println("os.Open(temp)")
	//	log.Fatal(err)
	//	return false
	//}
	config, err := png.DecodeConfig(bytes.NewBuffer(file_origin))
	if err != nil {
		fmt.Println("png.DecodeConfig(temp)")
		return false
	}
	/** 做等比缩放 */
	width := uint(200) /** 基准 */
	height := uint(200 * config.Height / config.Width)

	canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(toFilePath)
	defer file_out.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = png.Encode(file_out, canvas)
	if err != nil {
		fmt.Println("压缩图片失败")
		return false
	}
	return true
}
