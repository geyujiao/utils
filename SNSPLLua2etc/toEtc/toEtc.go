package snspllua2etc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
1. 按行读文件，得到list
2. list ----> map ，key=域名
做法：

	   1.按/CCS分割，或者按/SNS分割，（CCS/ats/remap.config）
	   2.区分边缘上层
	   3.CCS/ats 替换为 etc/trafficserver
	     SNS/ats 替换为 etc/trafficserver
		 CCS/nginx 替换为 etc/nginx
	     SNS/nginx 替换为 etc/nginx
		 域名.conf  替换为 /etc/nginx/sites-enabled/域名.conf
		 域名-l2.conf  替换为 /etc/nginx/sites-enabled/域名-l2.conf
	   4.写表格
*/

func OriginalToEtc(fileName string) (domainMapList1, domainMapList2 map[string][]string) {
	originalList := ReadOriginalData(fileName)
	domainMapList1, domainMapList2 = GetDomainMap(originalList)

	return
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

// 2. list ----> map ，key=域名
func GetDomainMap(originalList []string) (domainMapList1, domainMapList2 map[string][]string) {
	domainMapList1 = make(map[string][]string, 0)
	for _, original := range originalList {
		if strings.Contains(original, "SNS") {
			domain := strings.Split(original, "/")[1]
			if _, exist := domainMapList1[domain]; !exist {
				domainMapList1[domain] = make([]string, 0)
			}
			info := strings.Split(original, "SNS")[1]
			info = ToEtc(domain, info)
			domainMapList1[domain] = append(domainMapList1[domain], info)
		}
	}
	domainMapList2 = make(map[string][]string, 0)
	for _, original := range originalList {
		if strings.Contains(original, "CCS") {
			domain := strings.Split(original, "/")[1]
			if _, exist := domainMapList1[domain]; !exist {
				domainMapList2[domain] = make([]string, 0)
			}
			info := strings.Split(original, "CCS")[1]
			info = ToEtc(domain, info)
			domainMapList2[domain] = append(domainMapList2[domain], info)
		}
	}

	return
}

// 3. 替换
func ToEtc(domain, oriStr string) (newStr string) {
	///ats/remap.config
	newStr = strings.Replace(oriStr, "ats", "etc/trafficserver", 1)
	newStr = strings.Replace(newStr, "nginx", "etc/nginx", 1)
	if strings.Contains(oriStr, "/"+domain+".conf") {
		newStr = strings.Replace(newStr, "nginx", "etc/nginx/sites-enabled", 1)
	} else if strings.Contains(oriStr, "/"+domain+"-l2.conf") {
		newStr = strings.Replace(newStr, "nginx", "etc/nginx/sites-enabled", 1)
	}
	return newStr
}
