package toutils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func HeaderList(fileName string) (headerMap map[string]int) {
	contentList := ReadOriginalData(fileName)
	headerMap = make(map[string]int, 0)
	for _, info := range contentList {
		if strings.Contains(info, "header_filter_by_lua_file") {
			strs := strings.Split(info, "header_filter_by_lua_file")
			key := strings.Replace(strs[1], " ", "", -1)
			headerMap[key] = headerMap[key] + 1
		}

	}
	return headerMap
}

func UserAgentList(fileName, newFileName string) (resultList []string) {
	contentList := ReadOriginalData(fileName)
	for _, info := range contentList {
		if strings.Contains(info, "L1/") || strings.Contains(info, "L2/") {

			resultList = append(resultList, info)
		}

	}
	writerFile(newFileName, resultList)

	return resultList
}
func LogList(fileName string) (headerMap map[string]int) {
	contentList := ReadOriginalData(fileName)
	headerMap = make(map[string]int, 0)
	for _, info := range contentList {
		if strings.Contains(info, "log_by_lua_file") {
			strs := strings.Split(info, "log_by_lua_file")
			key := strings.Replace(strs[1], " ", "", -1)
			headerMap[key] = headerMap[key] + 1
		}

	}
	return headerMap
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

func writerFile(filename string, lines []string) {

	// 创建/打开文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // 确保文件在函数结束时关闭

	// 创建一个*Writer，用于按行写入
	writer := bufio.NewWriter(file)

	// 按行写入数据
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// 确保所有数据都被刷新到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// fmt.Println("文件写入成功，内容如下：")
	// 打印文件内容
	// content, _ := os.ReadFile(filename)
	// fmt.Print(string(content))
}
