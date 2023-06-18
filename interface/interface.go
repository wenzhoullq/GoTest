package main

type A interface {
	B
	C
}

type B interface{ B1() }

type C interface{ C1() }
type s struct{}

func (s s) C1() {}

func (s s) B1() {}

// 必须要实现所有的A的接口,如果有一个接口没有实现,那么则报错;这种写法用于大量的接口且保证所有的接口都被实现
func main() {
	var s s
	var a A = s
	a.B1()
}
