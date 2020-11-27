package chars

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//ToInt 转换成int格式
func ToInt(num interface{}) int {
	switch num.(type) {
	default:
		logger.Error("invalid type ", num)
		return 0

	case string:
		str := num.(string)
		index := strings.Index(str, ".")
		if index != -1 {
			str = str[:index]
		}

		result, _ := strconv.Atoi(str)
		return result

	case int:
		return num.(int)

	case int32:
		return int(num.(int32))

	case int64:
		return int(num.(int64))

	case float64:
		return int(num.(float64))

	case bool:
		if num.(bool) {
			return 1
		}
		return 0

	case []byte:
		result := 0
		for k, v := range num.([]byte) {
			result += (int(v) - 48) * int(math.Pow10(len(num.([]byte))-k-1))
		}
		return result

	case nil:
		return 0

	}
}

//ToFloat64 转换成float64格式
func ToFloat64(num interface{}) float64 {
	switch num.(type) {
	default:
		logger.Error("invalid type ", num)
		return 0

	case float64:
		return num.(float64)

	case string:
		result, _ := strconv.ParseFloat(num.(string), 64)
		return result

	case int:
		return float64(num.(int))

	case int32:
		return float64(num.(int32))

	case int64:
		return float64(num.(int64))

	case nil:
		return 0

	}
}

//ToString 转换成string
func ToString(num interface{}, prec ...int) string {
	var p = 2
	if len(prec) != 0 {
		p = prec[0]
	}

	switch num.(type) {
	default:
		logger.Error("invalid type ", num)
		return ""

	case string:
		str := num.(string)
		if len(prec) == 0 {
			return str
		}

		str = strings.Trim(str, "0")
		index := strings.Index(str, ".")
		if index == -1 {
			str += "."
			index = len(str) - 1
		}

		str += addZero(p)

		str = str[:index+p+1]

		if str == "" || str[0] == '.' {
			str = "0" + str
		}

		if str[len(str)-1] == '.' {
			str = str[:len(str)-1]
		}

		return str

	case int:
		return strconv.Itoa(num.(int))

	case int32:
		return strconv.Itoa(int(num.(int32)))

	case int64:
		return strconv.FormatInt(num.(int64), 10)

	case float64:
		return strconv.FormatFloat(num.(float64), 'f', p, 64)

	case nil:
		return ""

	}
}


func addZero(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "0"
	}

	return result
}

// float保留两位小数
func FloatRetain(f float64) float64  {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)

	return value
}


// 生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}