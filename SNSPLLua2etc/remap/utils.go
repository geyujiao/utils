package remap

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetMap(){
	list := ReadOriginalData("remap.config")
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	for _, info := range list {
		if strings.Contains(info, "cachepolicy.so") {
			map1[info] = "1"
		}
		info2 := strings.ReplaceAll(info, " ", "")
		if strings.Contains(info2, "@plugin=cachepolicy.so@pparam") {
			map2[info] = "1"
		}
	}
	result1 := make(map[string]string)
	result2 := make(map[string]string)

	for k1,_ := range map1 {
		if map2[k1] != "1" {
			result1[k1] = "1"
			fmt.Println("1111---" + k1)
		}
	}
	for k2,_ := range map2 {
		if map1[k2] != "1" {
			result2[k2] = "1"
			fmt.Println("2222---" + k2)
		}
	}
}


// 1. 按行读文件，得到list
func ReadOriginalData(fileName string) (originalList []string) {
	// 读取一个文件的内容
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	// 处理结束后关闭文件
	defer file.Close()

	// 使用bufio读取
	r := bufio.NewReader(file)

	for {
		// 分行读取文件  ReadLine返回单个行，不包括行尾字节(\n  或 \r\n)
		data, _, err := r.ReadLine()
		// 读取到末尾退出
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}
		if string(data) == "" {
			continue
		}
		originalList = append(originalList, string(data))
	}
	return
}
