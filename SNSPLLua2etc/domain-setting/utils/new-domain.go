package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
abuy.hihonor.com
aicloud-obs-auth.op.hihonorcloud.com
aicloud-obs.op.hihonorcloud.com
*/

func NewDomainList(srcHost string, destHostList []string) {
	for _, info := range destHostList {
		go NewDomain(srcHost, info)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Replacement completed successfully.")
}

func NewDomain(srcHost, destHost string) {

	err := copyDir("./"+srcHost, "./"+destHost)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	time.Sleep(2 * time.Second)

	// 修改文件名
	// ModifyFileName(fmt.Sprintf("%s/CCSL3/nginx/%s3-l2.conf", destHost, srcHost), fmt.Sprintf("%s/CCSL3/nginx/%s3-l2.conf", destHost, destHost))
	ModifyFileName(fmt.Sprintf("%s/CCS/nginx/%s-l2.conf", destHost, srcHost), fmt.Sprintf("%s/CCS/nginx/%s-l2.conf", destHost, destHost))
	ModifyFileName(fmt.Sprintf("%s/SNS/nginx/%s.conf", destHost, srcHost), fmt.Sprintf("%s/SNS/nginx/%s.conf", destHost, destHost))
	time.Sleep(5 * time.Second)

	fileList := []string{
		// fmt.Sprintf("%s/ccs/ats/parent.config", destHost),
		// fmt.Sprintf("%s/CCSL3/ats/remap.config", destHost),
		// fmt.Sprintf("%s/CCSL3/nginx/%s3-l2.conf", destHost, destHost),
		fmt.Sprintf("%s/CCS/ats/remap.config", destHost),
		fmt.Sprintf("%s/CCS/nginx/%s-l2.conf", destHost, destHost),
		fmt.Sprintf("%s/SNS/ats/parent.config", destHost),
		fmt.Sprintf("%s/SNS/ats/remap.config", destHost),
		fmt.Sprintf("%s/SNS/nginx/%s.conf", destHost, destHost),
	}
	// 修改文件内容
	for _, filePath := range fileList {
		err = replaceContentInFile(filePath, srcHost, destHost)
		if err != nil {
			fmt.Println("filePath:"+filePath+", Error:", err)
			os.Exit(1)
		}
	}

}

func replaceContentInFile(filePath, oldStr, newStr string) error {
	// 读取文件内容
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 替换内容
	output := strings.ReplaceAll(string(input), oldStr, newStr)

	// 写入新内容到文件
	err = ioutil.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

// 修改文件名
func ModifyFileName(src, dest string) {
	// 检查旧文件是否存在
	if _, err := os.Stat(src); os.IsNotExist(err) {
		fmt.Printf("文件 %s 不存在\n", src)
		return
	}

	// 重命名文件
	err := os.Rename(src, dest)
	if err != nil {
		fmt.Printf("无法重命名文件: %s\n", err)
		return
	}

	fmt.Printf("文件已成功重命名为 %s\n", dest)
}

func copyDir(src string, dest string) (err error) {
	// 创建目标文件夹
	if err = os.MkdirAll(dest, os.ModePerm); err != nil {
		return err
	}

	// 读取源文件夹内容
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		// 判断是否是文件夹
		if entry.IsDir() {
			// 递归复制文件夹
			if err = copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			// 复制文件
			if err = copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return
}

func copyFile(src, dest string) (err error) {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 复制文件内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return
}
