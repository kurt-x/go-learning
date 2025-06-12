package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 循环，go 只有 for 一种循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}

	// for 也是 while
	for sum < 1000 {
		sum += sum
	}

	// 没有条件时等同于 while(true) 死循环
	for {
		break
	}

	// if
	if true {
	} else if true {
	} else {
	}

	// if 前面可以加一个语句
	//goland:noinspection GoBoolExpressions
	if i := 0; i < 10 {
	}

	// switch
	// 顺序执行
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	case os + "test":
		fmt.Println("test")
	default:
		fmt.Println("Windows")
	}

	// 无条件 switch
	// 可用于替代 if-else-if
	switch {
	case true:
	case false:
	}
}
