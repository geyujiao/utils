package chars

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
	"strconv"
	"testing"
	"time"
)

func TestToInt(t *testing.T) {
	param := map[string]interface{}{}
	param["int10"] = int(10)
	param["int6410"] = int64(10)

	param["float10.0"] = 10.0
	param["float10.02"] = 10.02
	param["float0.1"] = 0.1
	param["float0.01"] = 0.01
	param["float0.10"] = 0.10

	param["string"] = ""
	param["string0"] = "0"
	param["string10"] = "10"
	param["string10.0"] = "10.0"
	param["string10.01"] = "10.01"
	param["string0.1"] = "0.1"
	param["string0.01"] = "0.01"
	param["string0.10"] = "0.10"

	for k, v := range param {
		logger.Info(k, "ToInt", ToInt(v))
		logger.Info(k, "ToFloat64", ToFloat64(v))
		logger.Info(k, "ToString", ToString(v))
	}
}

func TestFloatRetain(t *testing.T) {
	num := 0.123456

	value1, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", num), 64)
	logger.Info("11", value1)

	value2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	logger.Info("22", value2)

	value3, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", num), 64)
	logger.Info("33", value3)

	// 测试不通过
	//value4, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
	//logger.Info("44", value4)

}

func TestToString(t *testing.T) {
	// string to int
	count, _ := strconv.Atoi("10")
	logger.Info("string to int", count)

	// string to float64
	f, _ := strconv.ParseFloat("10.00", 64)
	logger.Info("string to float64", f)

	// int to string
	str := strconv.Itoa(10)
	logger.Info("int to string", str)

	logger.Info("int to string", fmt.Sprintf("%v----%d", 10, 10)) //%v 和 %d 都可以

	// float to string
	logger.Info("float to string", fmt.Sprintf("%.2f", 0.12))

}

func TestToString2(t *testing.T) {
	now := time.Now()
	logger.Info("time now", now) // cst 中国沿海时间(北京时间)
	strTime := "2020-11-26 11:26:00"

	// string to time
	timeT, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local) // cst 中国沿海时间(北京时间)
	logger.Info("string to time", timeT)

	timeT2, _ := time.Parse("2006-01-02 15:04:05", strTime) // utc 世界协调时间(不要用这个，容易出错)
	logger.Info("string to time", timeT2)

	// time to string
	logger.Info("time to string", now.Format("2006-01-02 15:04:05"))
}
