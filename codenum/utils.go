package codenum

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func countLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func countLinesInDirectory(dirPath string) (int, error) {
	totalLines := 0
	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() { // 确保我们只处理文件
			lines, err := countLinesInFile(path)
			if err != nil {
				return err // 如果打开或读取文件时出错，则返回错误
			}
			totalLines += lines
		}
		return nil
	})
	if err != nil {
		return 0, err // 如果遍历目录时出错，则返回错误
	}
	return totalLines, nil
}

func GetCodeNum(directoryPath string) {
	// directoryPath := "./your-directory-path" // 替换为你的目录路径
	lines, err := countLinesInDirectory(directoryPath)
	if err != nil {
		fmt.Printf("Error counting lines: %v\n", err)
		return
	}
	fmt.Printf("Total lines in directory: %d\n", lines)
}
