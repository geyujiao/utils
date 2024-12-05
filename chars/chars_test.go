package chars

import (
	"fmt"
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
		fmt.Println(k, "ToInt", ToInt(v))
		fmt.Println(k, "ToFloat64", ToFloat64(v))
		fmt.Println(k, "ToString", ToString(v))
	}
}

func TestFloatRetain(t *testing.T) {
	num := 0.123456

	value1, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", num), 64)
	fmt.Println("11", value1)

	value2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	fmt.Println("22", value2)

	value3, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", num), 64)
	fmt.Println("33", value3)

	// 测试不通过
	//value4, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
	//fmt.Println("44", value4)

}

func TestToString(t *testing.T) {
	// string to int
	count, _ := strconv.Atoi("10")
	fmt.Println("string to int", count)

	// string to float64
	f, _ := strconv.ParseFloat("10.00", 64)
	fmt.Println("string to float64", f)

	// int to string
	str := strconv.Itoa(10)
	fmt.Println("int to string", str)

	fmt.Println("int to string", fmt.Sprintf("%v----%d", 10, 10)) //%v 和 %d 都可以

	// float to string
	fmt.Println("float to string", fmt.Sprintf("%.2f", 0.12))

}

func TestToString2(t *testing.T) {
	now := time.Now()
	fmt.Println("time now", now) // cst 中国沿海时间(北京时间)
	strTime := "2020-11-26 11:26:00"

	// string to time
	timeT, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local) // cst 中国沿海时间(北京时间)
	fmt.Println("string to time", timeT)

	timeT2, _ := time.Parse("2006-01-02 15:04:05", strTime) // utc 世界协调时间(不要用这个，容易出错)
	fmt.Println("string to time", timeT2)

	// time to string
	fmt.Println("time to string", now.Format("2006-01-02 15:04:05"))
}
