package main

import (
	"fmt"
	"time"
)

func f0() {
	//i变化的时候都是值复制,因此产生了闭包,且go还没执行的时候就输出了,严格来说这并不是闭包
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()

	}
}
func f1() {
	//如果在加入了time.Sleep,那么i不会变得这么快,因此输出值为1,2,3...
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Second)
	}
}

func f2() {
	//通过传参生成临时变量,i对x无影响
	for i := 0; i < 10; i++ {
		go func(x int) {
			fmt.Println(x, &x)
		}(i)
	}
}

func f3() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	//在低版本的go中,range 的v是一种值复制,严格来说这不是闭包
	for _, v := range arr {
		go func() {
			fmt.Println(v)
		}()
	}
}

func f4() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	//让go执行完休眠一会输出值为1,2,3,4,5,6...
	for _, v := range arr {
		go func() {
			fmt.Println(v, &v)
		}()
		time.Sleep(time.Second)
	}
}

func Add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func f5() {
	//这种才是真闭包,局部变量引用全局变量局部变量导致未释放
	tmp1 := Add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := Add(10)
	fmt.Println(tmp2(10), tmp2(20))
}

func f6() int {
	//return不是一个原子操作,是赋值,执行defer,返回,因为这里是一种匿名返回,i的值赋值一个匿名变量s,然后defer执行i的自增和打印,最后返回的是s
	i := 0
	defer func() {
		i++
		fmt.Println("defer1 print", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2 print", i)
	}()
	return i
}

func f7() (i int) {
	//这里返回值是有名变量,return返回和defer操作的都是同一个值
	i = 1
	defer func() {
		i++
		fmt.Println("defer1 print", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2 print", i)
	}()
	return i
}

func f8() {
	//因为作用域的缘故,defer1的i是一个新的局部变量,对i的改变是不能影响到局部变量里的i
	i := 0
	defer func() {
		fmt.Println("defer1 print", i, &i)
	}()
	defer fmt.Println("defer2 print", i, &i)
	for ; i < 10; i++ {

	}
	fmt.Println("print i", i)
}

func main() {
	//f0()
	//f1()
	//f2()
	//f4()
	f5()
	//fmt.Println("print value", f6())
	//fmt.Println("print value", f7())
	//f8()
	time.Sleep(time.Second)
}
