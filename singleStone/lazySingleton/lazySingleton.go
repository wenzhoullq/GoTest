package main

import (
	"fmt"
	"sync"
)

// 饥饿模式,使用sync.Once包解决并发问题
var lazySingleton *LazySingleton
var Once sync.Once

type LazySingleton struct {
}

func (l LazySingleton) print() {
	fmt.Println("lazySingleton")
}
func GetLazySingleton() *LazySingleton {
	Once.Do(func() {
		lazySingleton = &LazySingleton{}
	})
	return lazySingleton
}

func main() {
	singleton := GetLazySingleton()
	singleton.print()
}
