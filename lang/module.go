package main

import "go-learning/lib" // 导入包

// 导入多个包
import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	lib.Test()
}
