package main

import (
	"fmt"
	"time"
)

// 无缓冲轮询打印
func main() {
	channel1, channel2 := make(chan struct{}), make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("cat", i+1)
			channel2 <- struct{}{}
			<-channel1
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			<-channel2
			fmt.Println("dog", i+1)
			channel1 <- struct{}{}
		}
	}()
	time.Sleep(time.Second)
}
