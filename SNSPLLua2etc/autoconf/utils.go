package autoconf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Build() string {
	domainList := []string{
		"finderks.video.qq.com.03.cdnhwcqir15.com",
		"finderqv.video.qq.com.97.cdnhwcqir15.com",
		"finderbk.video.qq.com.30.cdnhwcqir15.com",
		"finderdma.video.qq.com.15.cdnhwcqir15.com",
		"finderd1p.video.qq.com.07.cdnhwcqir15.com",
		"finderdsp.video.qq.com.73.cdnhwcqir15.com",
		"finderhy.video.qq.com.20.cdnhwcqir15.com",
		"finderv1.video.qq.com.11.cdnhwcqir15.com",
	}
	str := `"./tencent-video/finderydy.video.qq.com/CCS/ats/parent.config /腾讯视频号/finderydy.video.qq.com/L2/etc/trafficserver/"
    "./tencent-video/finderydy.video.qq.com/CCS/ats/remap.config /腾讯视频号/finderydy.video.qq.com/L2/etc/trafficserver/"
    "./tencent-video/finderydy.video.qq.com/CCS/nginx/finderydy.video.qq.com-l2.conf  /腾讯视频号/finderydy.video.qq.com/L2/etc/nginx/sites-enabled/"
    "./tencent-video/finderydy.video.qq.com/SNS/ats/parent.config /腾讯视频号/finderydy.video.qq.com/L1/etc/trafficserver/"
    "./tencent-video/finderydy.video.qq.com/SNS/ats/remap.config /腾讯视频号/finderydy.video.qq.com/L1/etc/trafficserver/"
    "./tencent-video/finderydy.video.qq.com/SNS/nginx/finderydy.video.qq.com.conf  /腾讯视频号/finderydy.video.qq.com/L1/etc/nginx/sites-enabled/"
    `

	originStr := "finderydy.video.qq.com"
	result := ""
	for _, domain := range domainList {

		r := strings.ReplaceAll(str, originStr, domain)
		result = result + r

	}
	return result
}

func GetAllLua(fileName, file2 string) {
	list := ReadOriginalData(fileName)
	luaList := make([]string, 0)
	luaMap := make(map[string]int, 0)
	for _, info := range list {
		if strings.Contains(info, "/etc/nginx/lua/") {
			strs := strings.Split(info, "/etc/nginx/lua/")
			lua := ""
			if len(strs) >= 2 {
				lua = "/etc/nginx/lua/" + strs[1]
				if strings.Contains(lua, ".conf") {
					continue
				}
				lua = strings.Replace(lua, ";", "", -1)
				lua = strings.Replace(lua, `"`, "", -1)
				lua = strings.Replace(lua, `'`, "", -1)
				lua = strings.Replace(lua, ` `, "", -1)
			}
			if lua != "" && luaMap[lua] == 0 {
				luaMap[lua] = 1
				luaList = append(luaList, lua)
			}
		}

	}
	luaList2 := make([]string, 0)
	for _, lua := range luaList {
		if luaMap[lua] == 1 {
			if strings.Contains(lua, ".conf") {
				continue
			}
			luaList2 = append(luaList2, lua)
		}
	}
	writerFileByline(file2, luaList2)
}

func GetNoExistLua(fileL1, fileL2, fileGit, file1, file2, file3, file4 string) {
	listL1 := ReadOriginalData(fileL1)
	listL2 := ReadOriginalData(fileL2)
	listGit := ReadOriginalData(fileGit)
	resultMapL1 := make(map[string]int, 0)
	resultMapL2 := make(map[string]int, 0)
	resultMapGit := make(map[string]int, 0)
	resultMapCommon := make(map[string]int, 0)
	resultListCommon := make([]string, 0)

	println(fmt.Sprintf("L1 ori=%v", len(listL1)))
	println(fmt.Sprintf("L2 ori=%v", len(listL2)))
	println(fmt.Sprintf("git ori=%v", len(listGit)))

	for _, info := range listL1 {
		resultMapL1[info] = 1
	}
	for _, info := range listL2 {
		resultMapL2[info] = 1
	}
	for _, info := range listGit {
		resultMapGit[info] = 1
	}
	l1Count := 0
	l2Count := 0
	for lua, _ := range resultMapGit {
		if resultMapL1[lua] == 1 {
			l1Count++
		}
	}
	for lua, _ := range resultMapGit {
		if resultMapL2[lua] == 1 {
			l2Count++
		}
	}
	println(fmt.Sprintf("L1 common=%v L2 common=%v", l1Count, l2Count))

	for lua, _ := range resultMapGit {
		if resultMapL1[lua] == 1 || resultMapL2[lua] == 1 {
			delete(resultMapL1, lua)
			delete(resultMapL2, lua)
			delete(resultMapGit, lua)
			if resultMapCommon[lua] == 0 {
				resultListCommon = append(resultListCommon, lua)
				resultMapCommon[lua] = 1
			}
		}
	}

	println(fmt.Sprintf("L1 res=%v", len(resultMapL1)))
	println(fmt.Sprintf("L2 res=%v", len(resultMapL2)))
	println(fmt.Sprintf("git res=%v", len(resultMapGit)))
	println(fmt.Sprintf("common res=%v", len(resultMapCommon)))
	r1 := make([]string, 0)
	r2 := make([]string, 0)
	r3 := make([]string, 0)
	for _, info := range listL1 {
		if resultMapL1[info] == 1 {
			r1 = append(r1, info)
		}
	}
	for _, info := range listL2 {
		if resultMapL2[info] == 1 {
			r2 = append(r2, info)
		}
	}
	for _, info := range listGit {
		if resultMapGit[info] == 1 {
			r3 = append(r3, info)
		}
	}

	writerFileByline(file1, r1)
	writerFileByline(file2, r2)
	writerFileByline(file3, r3)
	writerFileByline(file4, resultListCommon)

}

func GetNoExistLua2(fileL1, fileL2, fileGit, file1, file2, file3, file4 string) {
	listL1 := ReadOriginalData(fileL1)
	listL2 := ReadOriginalData(fileL2)
	listGit := ReadOriginalData(fileGit)
	resultMapL1 := make(map[string]int, 0)
	resultMapL2 := make(map[string]int, 0)
	resultMapGit := make(map[string]int, 0)
	resultMapCommon := make(map[string]int, 0)
	resultListCommon := make([]string, 0)

	for _, info := range listL1 {
		s := strings.Split(info, "/")
		resultMapL1[s[len(s)-1]] = 1
	}
	for _, info := range listL2 {
		s := strings.Split(info, "/")
		resultMapL2[s[len(s)-1]] = 1
	}
	for _, info := range listGit {
		s := strings.Split(info, "/")
		resultMapGit[s[len(s)-1]] = 1
	}
	for lua, _ := range resultMapGit {
		if _, ok := resultMapL1[lua]; ok {
			delete(resultMapL1, lua)
			delete(resultMapGit, lua)
			if resultMapCommon[lua] == 0 {
				resultListCommon = append(resultListCommon, lua)
				resultMapCommon[lua] = 1
			}
		}
	}
	for lua, _ := range resultMapGit {
		if _, ok := resultMapL2[lua]; ok {
			delete(resultMapL2, lua)
			delete(resultMapGit, lua)
			if resultMapCommon[lua] == 0 {
				resultListCommon = append(resultListCommon, lua)
				resultMapCommon[lua] = 1
			}
		}
	}

	println(len(resultMapL1))
	println(len(resultMapL2))
	println(len(resultMapGit))
	print(len(resultMapCommon))
	r1 := make([]string, 0)
	r2 := make([]string, 0)
	r3 := make([]string, 0)
	for _, info := range listL1 {
		s := strings.Split(info, "/")
		tmp := s[len(s)-1]
		if resultMapL1[tmp] == 1 {
			r1 = append(r1, info)
		}
	}
	for _, info := range listL2 {
		s := strings.Split(info, "/")
		tmp := s[len(s)-1]
		if resultMapL2[tmp] == 1 {
			r2 = append(r2, info)
		}
	}
	for _, info := range listGit {
		s := strings.Split(info, "/")
		tmp := s[len(s)-1]
		if resultMapGit[tmp] == 1 {
			r3 = append(r3, info)
		}
	}

	writerFileByline(file1, r1)
	writerFileByline(file2, r2)
	writerFileByline(file3, r3)
	writerFileByline(file4, resultListCommon)

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

func writerFileByline(filename string, lines []string) {

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

func writerFileByMap(filename string, lines map[string]int) {

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
	for line, _ := range lines {
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
