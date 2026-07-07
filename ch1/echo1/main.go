package main

import (
	"fmt"
	"os"
)

// os.Args 是 slice
// os.Args[0]，是命令本身的名字
// s[i] 获取单个元素
// s[m:n] 获取子序列
// := 是短变量声明（short variable declaration），注意：短变量声明只能在函数体中使用
// i := 1 等价于 var i int = 1 编译器会自动推导变量类型，所以不需要写 int。

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//[]int 表示"int 类型的切片"
	a := []int{1, 2, 3}
	//如果省略切片表达式的 m 或 n，会默认传入 0 或 len(s)
	fmt.Println(a[1:])

	var b []int
	fmt.Println(b == nil)

	c := []int(nil)
	fmt.Println(c == nil)
}
