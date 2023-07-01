package main

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
