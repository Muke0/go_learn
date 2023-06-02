package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// func HelloTom() string {
// 	return "Jerry"
// }
// func TestHelloTom(t *testing.T) {
// 	output := HelloTom()
// 	expectOutput := "Tom"
// 	if output != expectOutput {
// 		t.Errorf("Expexted %s do not match actual %s", expectOutput, output)
// 	}
// }

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}

func TestJudgePassLineFalse(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, false, isPass)
}

func TestMain(m *testing.M) {
	//测试前：数据装载、配置初始化等前置工作
	code := m.Run()
	//测试后：释放资源等收尾工作
	os.Exit(code)
}
