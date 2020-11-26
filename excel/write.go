package excel

import (
	"github.com/tealeg/xlsx"
	"github.com/vgmdj/utils/logger"
)

type RowInfo struct {
	Value1 string
	Value2 string
}

func WriteExcel(fileName string, listData []RowInfo)(err error)  {
	// 新建文件和sheet
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	// 写文件
	for _, info := range listData {
		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = info.Value1
		ageCell := row.AddCell()
		ageCell.Value = info.Value2
	}
	err = file.Save(fileName)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("write file success")
	return nil
}


func WriteExcel2(fileName string, list [][]string)(err error)  {
	// 新建文件和sheet
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	// 写文件
	for _, info := range list {
		row := sheet.AddRow()
		for _, v := range info {
			nameCell := row.AddCell()
			nameCell.Value = v
		}
	}
	err = file.Save(fileName)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("write file success")
	return nil
}