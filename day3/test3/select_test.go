package main

import (
	"os"
	"testing"
)

func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select()
	}
}

func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}

func TestMain(m *testing.M) {
	//测试前：数据装载、配置初始化等前置工作
	code := m.Run()
	//测试后：释放资源等收尾工作
	os.Exit(code)
}
