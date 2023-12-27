package exist

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Domain() {
	//1 读文件
	//2 210个域名
	//3 210是否在文件里

	cache210 := ReadOriginalData("cache210t.txt")
	domainAll := ReadOriginalData("hy600.txt")
	domain108 := ReadOriginalData("hy108.txt")

	cache210Map := make(map[string]string, 0)
	for _, info := range cache210 {
		cache210Map[info] = info
	}
	fmt.Println("cache210Map len", len(cache210Map))
	hy600List := []string{}
	lackCacheList := []string{}
	for _, cache := range cache210 {
		exist := false
		for _, domain := range domainAll {
			if cache == domain {
				exist = true
				hy600List = append(hy600List, domain)
				break
			}
		}
		if !exist {
			lackCacheList = append(lackCacheList, cache)
		}
	}
	fmt.Println("hy600List len", len(hy600List))
	WriteTxt("hy600List.txt", hy600List)

	fmt.Println("lackCacheList len", len(lackCacheList))

	hy108result := []string{}
	lackCacheList2 := []string{}
	for _, cache := range lackCacheList {
		exist := false
		for _, domain := range domain108 {
			if cache == domain {
				exist = true
				hy108result = append(hy108result, domain)
				break
			}
		}
		if !exist {
			lackCacheList2 = append(lackCacheList2, cache)
		}
	}
	fmt.Println("hy108result len", len(hy108result))
	WriteTxt("hy108result.txt", hy108result)

	fmt.Println("lackCacheList2 len", len(lackCacheList2))
	WriteTxt("lackCacheList2.txt", lackCacheList2)

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

func WriteTxt(fileName string, bodyList []string) {
	// 创建一个新文件，写入内容 5句 "hello,Gardon"
	// 1.打开文件 f:/abc.txt
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	// 计时关闭file句柄
	defer file.Close()
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for _, info := range bodyList {
		writer.WriteString(info + "\n")
	}
	// 因为weiter是带缓存的，因此在调用WriteString方法时，其实内容是先写入缓存的，所以需要调用Flush方法
	// 将缓存的数据真正写入 文件中，否则文件就没有数据
	writer.Flush()
}
