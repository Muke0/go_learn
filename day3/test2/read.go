package main

import (
	"bufio"
	"os"
	"strings"
)

func ReadFirstLine() string {
	open, err := os.Open("log") //从目录中按文件名读取文件内容
	defer open.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}
