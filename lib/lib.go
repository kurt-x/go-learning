package lib

import "fmt"

// Test
// 大写字母开头的成员会被导出（相当于 public）
func Test() {
	fmt.Println("Hello, module!")
}

// 小写字母开头的成员不会被导出（相当于 private）
//
//goland:noinspection GoUnusedFunction
func test() {}
