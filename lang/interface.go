package main

import (
	"fmt"
	"math"
)

// 方法
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }

// 接口
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser = Vertex{3, 4}
	fmt.Println(a.Abs())

	var _any interface{} = nil // 空接口可以是任意值，类似 Java 的 Object
	fmt.Println(_any)

	v := a.(Vertex) // 断言 a 是 Vertex 类型，并将值赋给 v，如果不是会引发 panic
	fmt.Println(v)
	t, ok := a.(fmt.State) // 判断断言是否成功，这种方式不会引发 panic，若不成功，t 赋值 nil，ok 赋值 false
	fmt.Println(t, ok)

	// 类型选择
	switch tp := a.(type) {
	case fmt.Stringer:
	case fmt.Formatter:
	default:
		fmt.Println("type select", tp)
	}
}
