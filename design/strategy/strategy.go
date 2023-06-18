package main

import "fmt"

type Strategy interface {
	Discount(float64) float64
}

type s1 struct{}

func (s s1) Discount(price float64) float64 {
	return price * 0.8
}

type s2 struct{}

func (s s2) Discount(price float64) float64 {
	return price - 100
}

type Buyer struct {
	strategy Strategy
}

func (b Buyer) SetStrategy(strategy Strategy) {
	b.strategy = strategy
}

func (b Buyer) Buy(num float64) float64 {
	return b.strategy.Discount(num)
}

// 抽象出策略类接口,然后写出具体的策略类;同时定义环境类,环境类有set方法和执行方法,根据不同的策略类作为参数最后的执行方法是不同的
// strategy是策略类,buyer是环境类
func main() {
	s2 := s2{}
	s1 := s1{}
	buyer := Buyer{}
	buyer.SetStrategy(s1)
	fmt.Println(s1.Discount(100))
	buyer.SetStrategy(s2)
	fmt.Println(s2.Discount(100))
}
