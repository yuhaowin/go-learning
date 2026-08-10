package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota // 美元 index = 0
	EUR                 // 欧元 index = 1
	GBP                 // 英镑 index = 2
	RMB                 // 人民币 index = 3
)

// 数组支持指定一个索引和对应值列表的方式初始化，初始化索引的顺序是无关紧要的,未指定初始值的元素将用零值初始化
func test() {
	symbol := [...]string{GBP: "￡", RMB: "￥", USD: "$", EUR: "€"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}

func main() {

	// 在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算

	// 数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	sha256.Sum256([]byte("x"))

	test()
}
