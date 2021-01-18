package main

import (
	"fmt"
)

func main() {
	forAll(100)
	forPar(1000)
	while(2000)
	forWhile()
	switchDemo(4, 2)
	deferDemo()
}

func deferDemo() {
	//延迟执行
	defer fmt.Printf("defer execute!\n")
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	//先执行
	fmt.Printf("first print!\n")
}

func switchDemo(a... int) {
	if a != nil {
		switch a[0] {
		case 1:
			fmt.Printf("this is one : %g\n", "1")
		case 2:
			fmt.Printf("this is two\n")
		case a[1] + 2:
			fmt.Printf("this is dynamic case\n")
		default:
			fmt.Printf("this is default\n")
		}
	}

}

func forAll(a int) {
	sum := 0
	for i := 0; i < a; i++ {
		sum+=i
	}
	fmt.Println(sum)
}

func forPar(a int) {
	sum := 0
	//可以省略前后调价，类似于while
	for ; sum < a; {
		sum++
	}
	fmt.Println(sum)
}

func while(a int) {
	sum := 1
	//相当于while语句
	for sum<a {
		sum+=2
	}
	fmt.Println(sum)
}

func forWhile() {
	sum := 1
	//死循环，也可以再for 后边添加true
	for {
		sum += 10
		if sum > 10000 {
			break
		}
	}
	fmt.Println(sum)
}

