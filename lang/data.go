package main

import "fmt"

func main() {
	var _ string                       // 声明变量，_ 表示忽略变量，后续不使用
	var _, _ bool                      // 声明多个变量
	_ = "Hello, world!"                // 赋值变量
	var _ = "Hello, world!"            // 声明并赋值，可以自动推断类型
	var _, _, _ = true, false, "hello" // 声明并赋值多个变量
	v := "hello"                       // 直接声明并赋值，只能用在局部作用域使用
	v1, v2, v3 := true, 1, "yes"
	//goland:noinspection GoBoolExpressions
	fmt.Println(v, v1, v2, v3)

	// 声明常量，常量不可以使用 :=
	const _ = 1

	// 分组声明
	var (
		_ = true
		_ = 1
	)

	// 数据类型
	// int、uint、uintptr 位长度跟随系统
	var _ int
	var _ int8
	var _ int16
	var _ int32
	var _ int64
	var _ uint
	var _ uint8
	var _ uint16
	var _ uint32
	var _ uint64
	var _ uintptr
	var _ byte // uint8 别名
	var _ rune // int32 别名，表示一个 Unicode 码位

	var _ float32
	var _ float64

	var _ complex64
	var _ complex128

	var _ bool

	var _ string

	// 没有赋值的变量会被自动赋零值
	var (
		e1 int    // 0
		e2 bool   // false
		e3 string // ""
	)

	fmt.Printf("零值：(%d) (%t), (%s)", e1, e2, e3)
	pi := 3.14
	var i = int(pi) // 类型转换，不同类型必须显式转换
	fmt.Println(pi, i)

	var _ = 10   // int
	var _ = 10.0 // float64
	var _ = 0.1i // complex128
}
