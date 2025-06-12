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

	// 指针
	// *T 是指向 T 类型的指针
	var p *int

	p = &i
	fmt.Println(p, *p)
	*p = 100
	fmt.Println(p, *p)

	// 结构体
	type Vertex struct {
		X int
		Y int
	}

	vtx := Vertex{1, 2}
	vtx.X = 4
	fmt.Println(vtx.X, vtx.Y)

	ptr := &vtx
	fmt.Println(ptr.X, ptr.Y) // 隐式解引用

	_ = Vertex{Y: 3}

	// 数组: [n]T
	var _ [2]string
	arr := [3]string{"a", "b", "c"}
	fmt.Println(arr[2])

	// 切片: []T
	var _ []string
	sli := []string{"a", "b", "c"}
	fmt.Println(sli[2])

	// 使用索引创建切片
	_ = arr[1:2]
	_ = sli[1:2]
	_ = arr[1:]
	_ = arr[:2]
	_ = arr[:] // 引用全部
	// 更改切片会修改被引用数组中的数据
	// 直接声明切片就是先声明一个数组然后引用
	fmt.Println(len(arr[1:2]), cap(arr[1:2])) // 长度和容量

	// make 函数创建切片
	arr2 := make([]int, 5)
	fmt.Println("make", len(arr2), cap(arr2))
	arr2 = make([]int, 0, 5)
	fmt.Println("make", len(arr2), cap(arr2))

	// 映射
	var _ map[int]int
	m := map[string]Vertex{"Bell Labs": {4, 4}}
	m["Google"] = Vertex{5, 6}
	m["Google"] = Vertex{7, 8}
	fmt.Println(m)
	delete(m, "Bell Labs")
	fmt.Println(m)
	e := m["Google"]
	fmt.Println("element", e)
	element, exist := m["Google"]
	fmt.Println("element", element, exist)
	m = make(map[string]Vertex, 5)

	// 函数（也是数据类型的一种）
	fun := func() { fmt.Println("function data") }
	fun()
}
