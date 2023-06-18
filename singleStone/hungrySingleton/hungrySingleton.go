package main

import "fmt"

// 饥饿模式 在init中完成加载,就没有所谓的并发问题了
var hungrySingleton HungrySingleton

type HungrySingleton struct {
}

func (h HungrySingleton) print() {
	fmt.Println("HungrySingleton")
}

func GetSingleton() HungrySingleton {
	return hungrySingleton
}

func init() {
	hungrySingleton = HungrySingleton{}
}

func main() {
	singleton := GetSingleton()
	singleton.print()
}
