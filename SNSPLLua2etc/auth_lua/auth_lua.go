package auth_lua

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

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

func FileReplaceAll(oldPath, newPath string) {

	// 打开需要被替换的文件
	oldFile, err := os.Open(oldPath)
	if err != nil {
		fmt.Println("Open file error:", err)
		return
	}
	defer oldFile.Close()

	// 打开用于替换的文件
	newFile, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error:", err)
		return
	}
	defer newFile.Close()

	// 重命名文件
	err = os.Rename(oldPath, oldPath+".bak")
	if err != nil {
		fmt.Println("Rename file error:", err)
		return
	}

	err = os.Rename(newPath, oldPath)
	if err != nil {
		fmt.Println("Rename file error:", err)
		return
	}

	// err = os.Remove(oldPath + ".bak")
	// if err != nil {
	// 	fmt.Println("Remove file error:", err)
	// }

	fmt.Println("File replaced successfully!")
}

func CopyFile(srcName, destName string) (int64, error) {
	//打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	//打开目标文件   可读可写可创建
	det, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer det.Close()
	return io.Copy(det, src)
}
