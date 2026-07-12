package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(apply(func(a int, b int) int {
		return a + b
	}, 3, 4))

	var sub func(int, int) int = func(a int, b int) int {
		return a - b
	}

	fmt.Println(apply(sub, 10, 2))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	// a、b 本身是指针（存的是地址），*a、*b 是解引用，表示"指针指向的那个值"。
	// - a → 类型 *int，值是某个地址，比如 0xc0000140a0
	// - *a → 对这个地址做解引用，取出/写入该地址上存的 int 值
	fmt.Println(a, b)
	fmt.Println(*a, *b)
	// a b 已经是指针类型了，&a &b 是取指针自己的地址
	fmt.Println(&a, &b)
	*a, *b = *b, *a
	//a, b = b, a
	//fmt.Println(a, b)
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d, %d\n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	return op(a, b)
}
