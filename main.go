package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现所有接口
type A interface {
	B
	C
}

type B interface{ B1() }

type C interface{ C1() }
type s struct{}

func (s s) C1() {}

func (s s) B1() {}
func main() {
	//单例模式
	//懒汉模式
	lazySingleston := GetLazySingleston()
	lazySingleston.print()
	lazySingleston = GetLazySingleston()
	//饿汉模式
	hungrySingleston := GetHungrySingleston()
	hungrySingleston.print()

	//实现所有接口
	var ss s
	var a A = ss
	a.C1()
	//策略类
	buyer := Buyer{}
	buyer.SetStray(new(S1))
	fmt.Println(buyer.Buy(100))
	buyer.SetStray(new(S2))
	fmt.Println(buyer.Buy(100))
	//工厂模式
	appleFactory := AppleFactory{}
	apple := appleFactory.GetFruit()
	apple.show()
	//轮询打印无缓冲
	c1, c2 := make(chan struct{}), make(chan struct{})
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("cat")
			c1 <- struct{}{}
			<-c2
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			<-c1
			fmt.Println("dog")
			c2 <- struct{}{}
		}
	}()
	time.Sleep(time.Second)
}

// 策略模式
type Stragy interface {
	Discount(float64) float64
}

type S1 struct {
}

func (s1 *S1) Discount(num float64) float64 {
	return float64(num * 0.8)
}

type S2 struct {
}

func (s2 *S2) Discount(num float64) float64 {
	return num - 100
}

type Buyer struct {
	Stragy Stragy
}

func (Buyer *Buyer) SetStray(stragy Stragy) {
	Buyer.Stragy = stragy
}

func (Buyer *Buyer) Buy(num float64) float64 {
	return Buyer.Stragy.Discount(num)
}

// 抽象工厂模式
type Fruit interface {
	show()
}

type AbstractFactory interface {
	GetFruit() Fruit
}

type apple struct {
}

func (apple *apple) show() {
	fmt.Println("this is apple")
}

type AppleFactory struct {
	AbstractFactory
}

func (AppleFactory *AppleFactory) GetFruit() Fruit {
	return &apple{}
}

// 饿汉模式

type HungrySingleston struct {
}

var hungrySingleston *HungrySingleston

func init() {
	fmt.Println("hungrySingleston init")
	hungrySingleston = &HungrySingleston{}
}
func GetHungrySingleston() *HungrySingleston {
	return hungrySingleston
}
func (HungrySingleston *HungrySingleston) print() {
	fmt.Println("i am hungrySingleston")
}

// 懒汉模式
type LazySingleston struct {
}

var lazySingleston *LazySingleston
var once sync.Once

func GetLazySingleston() *LazySingleston {
	once.Do(func() {
		fmt.Println("LazySingleston init")
		lazySingleston = &LazySingleston{}
	})
	return lazySingleston
}
func (LazySingleston *LazySingleston) print() {
	fmt.Println("i am lazySingleston")
}
