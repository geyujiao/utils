package excel

import (
	"strconv"
	"testing"
)

func TestWriteExcel2(t *testing.T) {
	data := [][]string{}
	for i:= 1; i<= 10 ; i++{
		row := []string{}
		for j:= 1; j<=5 ; j++{
			row = append(row, strconv.Itoa(i) + strconv.Itoa(j))
		}
		data = append(data, row)
	}
	WriteExcel2("test2.xlsx", data)
}
