package main

import (
	"fmt"
	"time"
)

// tiktok 节拍器
func tiktok() <-chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- 1
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return ch
}

func main() {
	ch := tiktok()
	// ch <- 1 // 尝试写入数据， 不能通过编译

	// 循环读取
	for {
		i := <-ch
		fmt.Println(i)
	}
}
