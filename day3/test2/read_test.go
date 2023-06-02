package main

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// func TestProcessFirstLine(t *testing.T) {
// 	firstLine := ProcessFirstLine
// 	assert.Equal(t, "line00", firstLine)
// }

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}

func TestMain(m *testing.M) {
	//测试前：数据装载、配置初始化等前置工作
	code := m.Run()
	//测试后：释放资源等收尾工作
	os.Exit(code)
}
