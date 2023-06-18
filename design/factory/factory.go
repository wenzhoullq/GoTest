package main

type Fruit interface {
	show()
}
type AbstractFactory interface {
	GetFruit() Fruit
}
type apple struct {
}

func (a apple) show() {
	println("apple")
}

type AppleFactory struct {
	AbstractFactory
}

func (a AppleFactory) GetFruit() apple {
	return apple{}
}

// 抽象工厂模式 主要把Fruit和AbstractFactory两个接口抽象出来,同时声明方法;然后构造相应的实体类继承这两个接口并且实现相应的方法
func main() {
	appleFactory := AppleFactory{}
	apple := appleFactory.GetFruit()
	apple.show()
}
