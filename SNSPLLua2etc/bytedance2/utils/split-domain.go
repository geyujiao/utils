package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*
abuy.hihonor.com
aicloud-obs-auth.op.hihonorcloud.com
aicloud-obs.op.hihonorcloud.com
*/

func SplitDomainList(srcHost string, destHostList []string) {
	for _, info := range destHostList {
		SplitDomain(srcHost, info)
	}
}

func SplitDomain(srcHost, destHost string) {

	err := copyDir("./"+srcHost, "./"+destHost)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	time.Sleep(2 * time.Second)

	// 修改文件名
	ModifyFileName(fmt.Sprintf("%s/CCS/nginx/%s-l2.conf", destHost, srcHost), fmt.Sprintf("%s/CCS/nginx/%s-l2.conf", destHost, destHost))
	ModifyFileName(fmt.Sprintf("%s/SNS/nginx/%s.conf", destHost, srcHost), fmt.Sprintf("%s/SNS/nginx/%s.conf", destHost, destHost))
	ModifyFileName(fmt.Sprintf("%s/CCSL3/nginx/%s3-l2.conf", destHost, srcHost), fmt.Sprintf("%s/CCSL3/nginx/%s3-l2.conf", destHost, destHost))

	time.Sleep(5 * time.Second)

	fileList := []string{
		fmt.Sprintf("%s/CCS/nginx/%s-l2.conf", destHost, destHost),
		fmt.Sprintf("%s/SNS/nginx/%s.conf", destHost, destHost),
	}
	// fileList2 := []string{
	// 	fmt.Sprintf("%s/CCS/ats/remap.config", destHost),
	// 	fmt.Sprintf("%s/SNS/ats/parent.config", destHost),
	// 	fmt.Sprintf("%s/SNS/ats/remap.config", destHost),
	// }

	// 修改文件内容
	for _, filePath := range fileList {
		// ModifyFileContent(srcHost, destHost, filePath)
		ModifyFileLine(filePath, destHost)
	}
	filePathL3 := fmt.Sprintf("%s/CCSL3/nginx/%s3-l2.conf", destHost, destHost)

	ModifyFileLine2(filePathL3, destHost)

	time.Sleep(5 * time.Second)
	// 修改文件内容
	ModifyFileContentAts(fmt.Sprintf("%s/CCS/ats/remap.config", srcHost), fmt.Sprintf("%s/CCS/ats/remap.config", destHost), destHost)
	ModifyFileContentAts(fmt.Sprintf("%s/SNS/ats/parent.config", srcHost), fmt.Sprintf("%s/SNS/ats/parent.config", destHost), destHost)
	ModifyFileContentAts(fmt.Sprintf("%s/SNS/ats/remap.config", srcHost), fmt.Sprintf("%s/SNS/ats/remap.config", destHost), destHost)

	time.Sleep(5 * time.Second)

	fmt.Println("Replacement completed successfully.")
}

// 修改文件内容
func ModifyFileContent(src, dest, fileName string) {

	in, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			fmt.Println("last line:", string(line))
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		// if string(line) == "" || string(line) == "}" || string(line) == "{" {
		// 	continue
		// }
		newLine := string(line)
		if strings.Contains(string(line), "server_name") {
			newLine = "  server_name           " + dest + ";"
		}
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
	}
	fmt.Println("FINISH!")
}

// 修改文件内容
func ModifyFileContentAts(fileNameSrc, fileNameDest, dest string) {

	in, err := os.Open(fileNameSrc)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(fileNameDest, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()
	out.Truncate(0)

	br := bufio.NewReader(in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			fmt.Println("last line:", string(line))
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := ""
		if strings.Contains(string(line), dest) {
			newLine = string(line)
			// fmt.Println("write to file newLine:", newLine)

			_, err = out.WriteString(newLine + "\n")
			if err != nil {
				fmt.Println("write to file fail:", err)
				os.Exit(-1)
			}
		}

	}
	fmt.Println("FINISH!")
}

func ModifyFileLine(filePath, dest string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	i := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "server_name") {
			newContent := "  server_name           " + dest + ";"
			lines = append(lines, newContent)
		} else {
			lines = append(lines, scanner.Text())
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	output, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	writer.Flush()
}

func ModifyFileLine2(filePath, dest string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	i := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "server_name") {
			newContent := "  server_name " + dest + ";"
			lines = append(lines, newContent)
		} else {
			lines = append(lines, scanner.Text())
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	output, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	writer.Flush()
}
