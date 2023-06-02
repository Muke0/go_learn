package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	println("hello goroutine :" + fmt.Sprint(i))
}
func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}

func CalSquare() {
	src := make(chan int)     //无缓冲队列
	dest := make(chan int, 3) //有缓冲队列
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i //将i发送到src这个channel中
		}
	}()
	go func() {
		defer close(dest)
		for i := range src { //获取src中的数字
			dest <- i * i
		}
	}()
	for i := range dest { //从dest中获得最终结果
		println(i)
	}
}

var (
	x    int64
	lock sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func Add() {
	x = 0
	for i := 0; i < 500; i++ { //在数量级很小时，有可能加不加锁结果相同
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("WithoutLock:", x)
	x = 0
	for i := 0; i < 500; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("WithLock:", x)
}

func ManyGoWait() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
func main() {
	// CalSquare()
	//Add()
	ManyGoWait()
}
