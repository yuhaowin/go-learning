package main

import "fmt"

// func(int) int 在这里表示是一个函数类型 （function type）
// 表示“接收一个 int 参数、返回一个 int 值的函数”。它描述的是函数的签名（参数类型和返回值类型），而不关心函数叫什么名字。

//   func adder() func(int) int {
//
//  - adder 是一个函数，它自己不接收参数
//  - adder 的返回值类型是 func(int) int，也就是说 adder() 返回的不是一个普通的值（比如 int、string），而是另一个函数

func adder() func(int) int {
	sum := 0 // 自由变量，闭包
	return func(v int) int {
		sum += v
		return sum
	}
}

// type iAdder func(int) (int, iAdder) 定义了一个具名函数类型，它的底层类型（underlying type）是：
//
//	接收一个 int，返回一个 (int, iAdder)
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := range 10 {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}
