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

type pathInfo struct {
	DomainOrPlugin string // 域名/插件
	AddOrModify    string // 新增/修改
	Scope          string // 业务组范围
	Plugin         string // 插件
	Issue          int    // 下发顺序
	RollBack       int    // 回滚顺序
}

func WritePublishExcel(requirementFile, pathFile, publishExcel string) {
	requirementList := GetRequirementsList(requirementFile)
	requirementMap, domainMap := GetDomainDateMap(requirementFile, pathFile)

	// 新建文件和sheet
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		return
	}
	//序号	域名	边缘下发文件名	边缘下发目标目录	上层下发文件名	上层下发插件目标目录
	row := sheet.AddRow()
	nameCell := row.AddCell()
	nameCell.Value = "需求描述"
	nameCell = row.AddCell()
	nameCell.Value = "域名"
	nameCell = row.AddCell()
	nameCell.Value = "域名/插件"
	nameCell = row.AddCell()
	nameCell.Value = "新增/修改"
	nameCell = row.AddCell()
	nameCell.Value = "业务组范围"
	nameCell = row.AddCell()
	nameCell.Value = "插件"
	nameCell = row.AddCell()
	nameCell.Value = "下发顺序"
	nameCell = row.AddCell()
	nameCell.Value = "回滚顺序"

	for _, requirement := range requirementList {
		for _, domain := range requirementMap[requirement[0]] {
			for k, info := range domainMap[domain] {
				row := sheet.AddRow()
				requirementTmp := ""
				domainTmp := ""
				if k == 0 {
					requirementTmp = requirement[1]
					domainTmp = domain
				}
				nameCell := row.AddCell()
				nameCell.Value = requirementTmp
				nameCell = row.AddCell()
				nameCell.Value = domainTmp

				nameCell = row.AddCell()
				nameCell.Value = info.DomainOrPlugin // 域名/插件
				nameCell = row.AddCell()
				nameCell.Value = info.AddOrModify // 新增/修改
				nameCell = row.AddCell()
				nameCell.Value = info.Scope // 业务组范围
				nameCell = row.AddCell()
				nameCell.Value = info.Plugin // 插件
				nameCell = row.AddCell()
				nameCell.Value = strconv.Itoa(info.Issue) // 下发顺序
				nameCell = row.AddCell()
				nameCell.Value = strconv.Itoa(info.RollBack) // 回滚顺序
			}

		}

	}

	err = file.Save(publishExcel)
	if err != nil {
		return
	}
	return
}

func WritePublishExcel2(requirementFile, pathFile, publishExcel string) {
	requirementList := GetRequirementsList(requirementFile)
	requirementMap, _ := GetDomainDateMap(requirementFile, pathFile)

	// 新建文件和sheet
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		return
	}
	row := sheet.AddRow()
	nameCell := row.AddCell()
	nameCell.Value = "域名"
	nameCell = row.AddCell()
	nameCell.Value = "需求描述"

	for _, requirement := range requirementList {
		for _, domain := range requirementMap[requirement[0]] {
			row := sheet.AddRow()
			requirementTmp := requirement[1]
			domainTmp := domain

			nameCell := row.AddCell()
			nameCell.Value = domainTmp
			nameCell = row.AddCell()
			nameCell.Value = requirementTmp

		}

	}

	err = file.Save(publishExcel)
	if err != nil {
		println(err.Error())
		return
	}
	return
}
func GetDomainDateMap(requirementFile, pathFile string) (requirementMap map[string][]string, domainMap map[string][]pathInfo) {
	pathList := GetOriginPathList(pathFile)
	requirementMap = make(map[string][]string) // key=requirement value=[doamin1,doamin2]
	domainMap = make(map[string][]pathInfo)    // key=domain value=[pathInfo1,pathInfo2]

	domainListTmp := make([]string, 0)
	for _, pathStr := range pathList {
		// ./logs-scheme/所有域名/L1+L2+L3/etc/nginx/lua/log_wrong_handle.lua
		strList := strings.Split(pathStr, "/")
		requirement := strList[1]
		domain := strList[2]

		if _, exist := requirementMap[requirement]; !exist {
			requirementMap[requirement] = make([]string, 0)
		}

		if _, exist := domainMap[domain]; !exist {
			domainMap[domain] = make([]pathInfo, 0)
			requirementMap[requirement] = append(requirementMap[requirement], domain)
			domainListTmp = append(domainListTmp, domain)
		}

		info := pathInfo{}
		// DomainOrPlugin 域名/插件 // /etc/nginx/sites-enabled/ remap.config parent.config
		if strings.Contains(pathStr, "/etc/nginx/sites-enabled/") || strings.Contains(pathStr, "remap.config") || strings.Contains(pathStr, "parent.config") {
			info.DomainOrPlugin = domain
		} else {
			info.DomainOrPlugin = "插件"
		}
		info.AddOrModify = "修改" // 新增/修改 // 默认=修改，待确认

		// scope 业务组范围
		if strList[3] == "L1+L2+L3" || strList[3] == "L1+L2" {
			info.Scope = "全网"
		} else if strList[3] == "L2+L3" {
			info.Scope = "内容中心节点L2+L3"
		} else if strList[3] == "L1" {
			info.Scope = "边缘节点L1"
		} else if strList[3] == "L2" {
			info.Scope = "内容中心节点L2"
		} else if strList[3] == "L3" {
			info.Scope = "内容中心节点L3"
		}
		// Plugin 插件
		strs := strings.Split(pathStr, "/etc/")
		info.Plugin = "/etc/" + strs[1]
		// Issue 下发顺序
		info.Issue = len(domainMap[domain]) + 1

		domainMap[domain] = append(domainMap[domain], info)
	}

	// RollBack  回滚顺序
	for _, domain := range domainListTmp {
		l := len(domainMap[domain])
		for i := 0; i < l; i++ {
			domainMap[domain][i].RollBack = l - i
		}
	}

	return
}

// 获取原始下发文件路径
func GetOriginPathList(fileName string) (list []string) {
	list = writeFileByline(fileName)
	return list
}

// 获取需求list
func GetRequirementsList(fileName string) (requireList [][]string) {
	list := writeFileByline(fileName)
	for _, info := range list {
		strs := strings.Split(info, " ")
		requireList = append(requireList, []string{strs[0], strs[1]})
	}
	return requireList
}

// 1. 按行读文件，得到list
func writeFileByline(fileName string) (datelist []string) {
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
		datelist = append(datelist, string(data))
	}
	return
}
