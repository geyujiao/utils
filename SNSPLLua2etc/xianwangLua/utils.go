package xianwanglua

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func ResultExcel(excelFile string) {
	// 新建文件和sheet
	file := xlsx.NewFile()

	WriteL1Ng(file)

	WriteL2Ng(file)

	WriteL1Ats(file)

	WriteL2Ats(file)

	err := file.Save(excelFile)
	if err != nil {
		fmt.Println("err---", err.Error())
		return
	}
	return
}

func WriteL1Ng(file *xlsx.File) *xlsx.File {
	sheet, err := file.AddSheet("L1-ng")
	if err != nil {
		fmt.Println("err---", err.Error())
		return file
	}
	luaStageMap, luaList := L1Nginx("L1-nginx.txt", "L1-nginx-result.txt")
	// 序号	域名	边缘下发文件名	边缘下发目标目录	上层下发文件名	上层下发插件目标目录
	row := sheet.AddRow()
	nameCell := row.AddCell()
	nameCell.Value = "lua文件"
	nameCell = row.AddCell()
	nameCell.Value = "阶段"
	nameCell = row.AddCell()
	nameCell.Value = "相关域名"
	nameCell = row.AddCell()
	nameCell.Value = "需求文档"

	// 写文件
	for _, lua := range luaList {
		row := sheet.AddRow()
		// lua文件
		nameCell := row.AddCell()
		nameCell.Value = lua
		// 阶段
		nameCell = row.AddCell()
		nameCell.Value = luaStageMap[lua]

	}
	return file
}

func WriteL2Ng(file *xlsx.File) *xlsx.File {
	luaStageMap2, luaList2 := L2Nginx("L2-nginx.txt", "L2-nginx-result.txt")
	sheet2, err := file.AddSheet("L2-ng")
	if err != nil {
		fmt.Println("err---", err.Error())
		return file
	}
	// 序号	域名	边缘下发文件名	边缘下发目标目录	上层下发文件名	上层下发插件目标目录
	row2 := sheet2.AddRow()
	nameCell2 := row2.AddCell()
	nameCell2.Value = "lua文件"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "阶段"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "相关域名"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "需求文档"

	// 写文件
	for _, lua := range luaList2 {
		row := sheet2.AddRow()
		// lua文件
		nameCell := row.AddCell()
		nameCell.Value = lua
		// 阶段
		nameCell = row.AddCell()
		nameCell.Value = luaStageMap2[lua]
	}
	return file
}

func WriteL1Ats(file *xlsx.File) *xlsx.File {
	luaMap, luaList := L1Ats("L1-ats.txt", "L1-ats-result.txt")
	sheet, err := file.AddSheet("L1-ats")
	if err != nil {
		fmt.Println("err---", err.Error())
		return file
	}
	// 序号	域名	边缘下发文件名	边缘下发目标目录	上层下发文件名	上层下发插件目标目录
	row2 := sheet.AddRow()
	nameCell2 := row2.AddCell()
	nameCell2.Value = "lua文件"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "相关域名"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "需求文档"

	// 写文件
	for _, lua := range luaList {
		row := sheet.AddRow()
		// lua文件
		nameCell := row.AddCell()
		nameCell.Value = luaMap[lua]
	}
	return file
}

func WriteL2Ats(file *xlsx.File) *xlsx.File {
	luaMap, luaList := L2Ats("L2-ats.txt", "L2-ats-result.txt")
	sheet, err := file.AddSheet("L2-ats")
	if err != nil {
		fmt.Println("err---", err.Error())
		return file
	}
	// 序号	域名	边缘下发文件名	边缘下发目标目录	上层下发文件名	上层下发插件目标目录
	row2 := sheet.AddRow()
	nameCell2 := row2.AddCell()
	nameCell2.Value = "lua文件"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "相关域名"
	nameCell2 = row2.AddCell()
	nameCell2.Value = "需求文档"

	// 写文件
	for _, lua := range luaList {
		row := sheet.AddRow()
		// lua文件
		nameCell := row.AddCell()
		nameCell.Value = luaMap[lua]
	}
	return file
}

func L1Nginx(fileName, fileName2 string) (luaStageMap map[string]string, luaList []string) {
	list := ReadOriginalData(fileName)
	luaMap := make(map[string]int, 0)
	luaStageMap = make(map[string]string, 0) // 阶段
	// luaDoaminMap := make(map[string][]string, 0) // 域名
	luaList = make([]string, 0)
	for _, info := range list {
		if strings.Contains(info, ".lua") {
			if strings.Contains(info, "#DEPEND_FILE") {
				continue
			}
			if strings.Contains(info, "#") {
				continue
			}
			strs := strings.Split(info, "/etc/nginx/")
			if len(strs) >= 2 {
				key := "/etc/nginx/" + strs[1][:len(strs[1])-1]
				key = strings.ReplaceAll(key, ";", "")
				key = strings.ReplaceAll(key, "'", "")
				key = strings.ReplaceAll(key, " ", "")
				key = strings.ReplaceAll(key, "\"", "")

				if _, exist := luaMap[key]; !exist {
					luaList = append(luaList, key)
				}

				luaMap[key] = luaMap[key] + 1

				str2 := strings.Split(strs[0], ":")
				if len(str2) == 2 {
					stage := strings.ReplaceAll(str2[1], "\"", "")
					stage = strings.ReplaceAll(stage, "'", "")
					luaStageMap[key] = stage

				}

			}
		}

	}
	// luaNumTotal := 0
	// for _, num := range luaMap {
	// 	luaNumTotal = luaNumTotal + num
	// }
	// // bytes, _ := json.Marshal(luaMap)
	// // writerFile("L1-nginx.json", bytes)

	// println("luaNumTotal-----=", luaNumTotal) //luaNumTotal-----= 2823
	// writerFileByline(fileName2, luaList)

	return luaStageMap, luaList
}

func L2Nginx(fileName, fileName2 string) (luaStageMap map[string]string, luaList []string) {
	list := ReadOriginalData(fileName)
	luaMap := make(map[string]int, 0)
	luaStageMap = make(map[string]string, 0) // 阶段
	// luaDoaminMap := make(map[string][]string, 0) // 域名
	luaList = make([]string, 0)
	for _, info := range list {
		if strings.Contains(info, ".lua") {
			if strings.Contains(info, "#DEPEND_FILE") {
				continue
			}
			if strings.Contains(info, "#") {
				continue
			}
			strs := strings.Split(info, "/etc/nginx/")
			if len(strs) >= 2 {
				key := "/etc/nginx/" + strs[1][:len(strs[1])-1]
				key = strings.ReplaceAll(key, ";", "")
				key = strings.ReplaceAll(key, "'", "")
				key = strings.ReplaceAll(key, " ", "")
				key = strings.ReplaceAll(key, "\"", "")

				if _, exist := luaMap[key]; !exist {
					luaList = append(luaList, key)
				}

				luaMap[key] = luaMap[key] + 1

				str2 := strings.Split(strs[0], ":")
				if len(str2) == 2 {
					stage := strings.ReplaceAll(str2[1], "\"", "")
					stage = strings.ReplaceAll(stage, "'", "")
					luaStageMap[key] = stage

				}

			}
		}

	}
	// luaNumTotal := 0
	// for _, num := range luaMap {
	// 	luaNumTotal = luaNumTotal + num
	// }
	// // bytes, _ := json.Marshal(luaMap)
	// // writerFile("L2-nginx.json", bytes)

	// println("luaNumTotal---L2--=", luaNumTotal) // luaNumTotal---L2--= 1572
	// writerFileByline(fileName2, luaList)

	return luaStageMap, luaList
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

func writerFile(fileName string, date []byte) {
	err := os.WriteFile(fileName, date, 0644)
	if err != nil {
		fmt.Println("writerFile err " + err.Error())
	}
}
