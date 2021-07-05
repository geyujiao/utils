package image

import (
	"github.com/vgmdj/utils/logger"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileType(t *testing.T) {
	filename := ""
	//filename = "gyj.mp4"
	//filename = "cat.png"
	//filename = "huoti.mp4"
	//filename = "hz2.jpeg"
	filename = "dlrb.jpg"
	//f, err := os.Open("C:\\Users\\Administrator\\Desktop\\api.html")
	f, err := os.Open(filename)
	//f, err := os.Open("cat.png")

	if err != nil {
		logger.Error("open error: ", err)
	}

	fSrc, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Error("open error: ", err)
	}

	logger.Info("file type", GetFileType(fSrc[:10]))
}
