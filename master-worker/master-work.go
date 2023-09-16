package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//master - work 工作模型 1、master负责分发计算任务，任务总数1000
//2、每一个计算任务由master调起新的work协程，master退出前，必须通知所有work先退出，自己再退出
//3、同时运行的work协程数不超过3个,work超时5s自动退出 4、work计算结果返回给master

func main() {
	type Cal func(int) int
	var cal Cal = func(i int) int {
		return i
	}
	ch := make(chan Cal, 3)
	ans := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		ch <- cal
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		out:
			for {
				select {
				case f := <-ch:
					time.Sleep(1 * time.Second) //用于模拟计算时间时间
					ans += f(i)
					break out //一定要加break,不然无法退出
				case <-ctx.Done():
					break out
				default:
				}
			}
		}(i)
		go func() {
			time.Sleep(time.Second * 5)
			cancel()
		}()
	}
	wg.Wait()
	fmt.Println(ans)

}
