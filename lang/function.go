package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// 可以简写，x、y 都是 int 类型
func add2(x, y int) int {
	return x + y
}

// 返回值可以被命名
func add3(x, y int) (r int) {
	r = x + y
	// return r
	return // 可以使用裸返回语句，返回被命名的返回值
}

// 多返回值
func hello() (string, string) {
	return "hello", "world"
}

// defer 会将语句推迟到当前函数返回后执行，函数调用的参数语句会被立即执行
func deferFunc() {
	defer fmt.Println("world!")

	fmt.Println("hello")

	// defer 调用会被压入一个栈中，按先进后出顺序执行
	for i := 0; i < 10; i++ {
		defer fmt.Println("defer", i)
	}
}

func main() {
	deferFunc()

	v1, v2 := hello()
	fmt.Println(v1, v2)
}
