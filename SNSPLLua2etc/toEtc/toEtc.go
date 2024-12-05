package snspllua2etc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
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

func OriginalToEtc(oriFileName, targFileName string) {
	originalList := ReadOriginalData(oriFileName)
	domainMapList1, domainMapList2 := GetDomainMap(originalList)
	domainMapTmp1, domainMapTmp2 := GetDomainMapTmp(originalList)
	WriteExcelEtc(targFileName, domainMapTmp1, domainMapTmp2, domainMapList1, domainMapList2)
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

// 2. list ----> map ，key=域名---下发路径/etc/nginx/sites-enabled/xx.conf
func GetDomainMap(originalList []string) (domainMapList1, domainMapList2 map[string][]string) {
	domainMapList1 = make(map[string][]string, 0)
	for _, original := range originalList {
		if strings.Contains(original, "SNS") { // migu/v6-sc.miguvideo.com/SNS/nginx/v6-sc.miguvideo.com.conf
			domain := strings.Split(original, "/")[1] // = v6-sc.miguvideo.com
			if _, exist := domainMapList1[domain]; !exist {
				domainMapList1[domain] = make([]string, 0)
			}
			info := strings.Split(original, "SNS")[1] //  = /nginx/v6-sc.miguvideo.com.conf
			info = ToEtc(domain, info)
			domainMapList1[domain] = append(domainMapList1[domain], info)
		}
	}
	domainMapList2 = make(map[string][]string, 0)
	for _, original := range originalList {
		if strings.Contains(original, "CCS/") {
			domain := strings.Split(original, "/")[1]
			if _, exist := domainMapList2[domain]; !exist {
				domainMapList2[domain] = make([]string, 0)
			}
			info := strings.Split(original, "CCS")[1]
			info = ToEtc(domain, info)
			domainMapList2[domain] = append(domainMapList2[domain], info)
		}
	}

	return
}

// 2. list ----> map ，key=域名----原始路径/SNS/nginx/xxx.conf
func GetDomainMapTmp(originalList []string) (domainMapTmp1, domainMapTmp2 map[string][]string) {
	domainMapTmp1 = make(map[string][]string, 0)
	for _, original := range originalList {
		if strings.Contains(original, "SNS") { // migu/v6-sc.miguvideo.com/SNS/nginx/v6-sc.miguvideo.com.conf
			domain := strings.Split(original, "/")[1] // = v6-sc.miguvideo.com
			if _, exist := domainMapTmp1[domain]; !exist {
				domainMapTmp1[domain] = make([]string, 0)
			}
			info := "/SNS" + strings.Split(original, "SNS")[1] //  = /SNS/nginx/v6-sc.miguvideo.com.conf
			domainMapTmp1[domain] = append(domainMapTmp1[domain], info)
		}
	}
	domainMapTmp2 = make(map[string][]string, 0)
	for _, original := range originalList {
		if strings.Contains(original, "CCS/") {
			domain := strings.Split(original, "/")[1]
			if _, exist := domainMapTmp2[domain]; !exist {
				domainMapTmp2[domain] = make([]string, 0)
			}
			info := "/CCS" + strings.Split(original, "CCS")[1]
			domainMapTmp2[domain] = append(domainMapTmp2[domain], info)
		}
	}

	return
}

// 3. 替换
func ToEtc(domain, oriStr string) (newStr string) { //  /nginx/v6-sc.miguvideo.com.conf
	///ats/remap.config
	newStr = strings.Replace(oriStr, "ats", "etc/trafficserver", 1)
	newStr = strings.Replace(newStr, "nginx", "etc/nginx", 1) //  /etc/nginx/v6-sc.miguvideo.com.conf
	if strings.Contains(oriStr, "/"+domain+".conf") {
		newStr = strings.Replace(newStr, "nginx", "nginx/sites-enabled", 1) // /etc/nginx/sites-enabled/v6-sc.miguvideo.com.conf
	} else if strings.Contains(oriStr, "/"+domain+"-l2.conf") {
		newStr = strings.Replace(newStr, "nginx", "nginx/sites-enabled", 1)
	}
	return newStr
}

// 4. 写文件
func WriteExcelEtc(fileName string, domainMapTmp1, domainMapTmp2, domainMapList1, domainMapList2 map[string][]string) (err error) {

	domainList := ReadOriginalData("snspllua-domain.txt")

	// 新建文件和sheet
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		return err
	}
	//序号	域名	边缘下发文件名	边缘下发目标目录	上层下发文件名	上层下发插件目标目录
	row := sheet.AddRow()
	nameCell := row.AddCell()
	nameCell.Value = "序号"
	nameCell = row.AddCell()
	nameCell.Value = "域名"
	nameCell = row.AddCell()
	nameCell.Value = "需求"
	nameCell = row.AddCell()
	nameCell.Value = "边缘下发文件名"
	nameCell = row.AddCell()
	nameCell.Value = "边缘下发目标目录"
	nameCell = row.AddCell()
	nameCell.Value = "上层下发文件名"
	nameCell = row.AddCell()
	nameCell.Value = "上层下发插件目标目录"

	// 写文件
	for k, domainAndContent := range domainList {
		str2 := strings.Split(domainAndContent, " ")
		domain := str2[0]
		row := sheet.AddRow()
		// 序号
		nameCell := row.AddCell()
		nameCell.Value = strconv.Itoa(k + 1)
		// 域名
		nameCell = row.AddCell()
		nameCell.Value = domain
		// 需求
		nameCell = row.AddCell()
		nameCell.Value = str2[1]
		// 边缘下发文件名
		value2 := ""
		for _, info := range domainMapTmp1[domain] {
			value2 = value2 + info + "\n"
		}
		nameCell = row.AddCell()
		nameCell.Value = value2
		// 边缘下发目标目录
		value1 := ""
		for _, info := range domainMapList1[domain] {
			value1 = value1 + info + "\n"
		}
		nameCell = row.AddCell()
		nameCell.Value = value1

		// 上层下发文件名
		if domainMapTmp2[domain] != nil {
			value2 := ""
			for _, info := range domainMapTmp2[domain] {
				value2 = value2 + info + "\n"
			}
			nameCell := row.AddCell()
			nameCell.Value = value2
		}
		// 上层下发插件目标目录
		if domainMapList2[domain] != nil {
			value2 := ""
			for _, info := range domainMapList2[domain] {
				value2 = value2 + info + "\n"
			}
			nameCell := row.AddCell()
			nameCell.Value = value2
		}

	}
	// for domain, _ := range domainMapTmp1 {
	// 	// str2 := strings.Split(domainAndContent, " ")
	// 	// domain := str2[0]
	// 	row := sheet.AddRow()
	// 	// 序号
	// 	nameCell := row.AddCell()
	// 	nameCell.Value = strconv.Itoa(0 + 1)
	// 	// 域名
	// 	nameCell = row.AddCell()
	// 	nameCell.Value = domain
	// 	// 需求
	// 	nameCell = row.AddCell()
	// 	nameCell.Value = ""
	// 	// 边缘下发文件名
	// 	if domainMapTmp1[domain] != nil {
	// 		value2 := ""
	// 		for _, info := range domainMapTmp1[domain] {
	// 			value2 = value2 + info + "\n"
	// 		}
	// 		nameCell := row.AddCell()
	// 		nameCell.Value = value2
	// 	}
	// 	// 边缘下发目标目录
	// 	value1 := ""
	// 	for _, info := range domainMapList1[domain] {
	// 		value1 = value1 + info + "\n"
	// 	}
	// 	nameCell = row.AddCell()
	// 	nameCell.Value = value1

	// 	// 上层下发文件名
	// 	if domainMapTmp2[domain] != nil {
	// 		value2 := ""
	// 		for _, info := range domainMapTmp2[domain] {
	// 			value2 = value2 + info + "\n"
	// 		}
	// 		nameCell := row.AddCell()
	// 		nameCell.Value = value2
	// 	}
	// 	// 上层下发插件目标目录
	// 	if domainMapList2[domain] != nil {
	// 		value2 := ""
	// 		for _, info := range domainMapList2[domain] {
	// 			value2 = value2 + info + "\n"
	// 		}
	// 		nameCell := row.AddCell()
	// 		nameCell.Value = value2
	// 	}

	// }
	err = file.Save(fileName)
	if err != nil {
		return err
	}
	return nil
}

func RefererList(fileName, newFileName string) {
	originalList := ReadOriginalData(fileName)
	referMap := make(map[string]int)
	for _, info := range originalList {
		if _, exist := referMap[info]; !exist {
			referMap[info] = 1
		} else {
			count := referMap[info]
			referMap[info] = count + 1
		}
	}

	newList := []string{}
	for refer, count := range referMap {
		if count > 1 {
			println(fmt.Sprintf("refer=%s, count=%v", refer, count))
		}
		newList = append(newList, refer)
	}
	WriteFile(newList, newFileName)

}

func WriteFile(lines []string, filename string) {

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
	// // 打印文件内容
	// content, _ := os.ReadFile(filename)
	// fmt.Print(string(content))
}
