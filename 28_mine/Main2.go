package main

import (
	"fmt"
	"time"
)

func main() {
	println("Hello, World!")

	// 创建一个无缓冲的channel
	ch := make(chan int)

	// 启动一个goroutine，向channel发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 向channel发送数据
			fmt.Println("Sent:", i)
		}
		close(ch) // 发送完毕后关闭channel
	}()

	// 从channel接收数据
	for v := range ch {
		fmt.Println("Received:", v)
		time.Sleep(1 * time.Second) // 模拟处理时间
	}

}
