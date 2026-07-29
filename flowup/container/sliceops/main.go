package main

import "fmt"

func main() {

	fmt.Println("Creating slice")

	// 使用var s []int 声明切片变量，此时切片值为 nil，底层数组未分配
	// Go 语言中切片的零值为 nil，但可以直接进行 append 操作而不会崩溃
	var s []int // Zero value for slice is nil

	fmt.Println(s == nil)
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := make([]int, 16)     // slice len
	s2 := make([]int, 10, 32) // slice len cap

	fmt.Println(s1, s2)

	fmt.Println("Copying slice")

	copy(s1, s) // copy s to s1
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println("Deleting elements from slice")

	s1 = append(s1[:3], s1[4:]...)
	fmt.Println(s1, len(s1), cap(s1))

}

func printSlice(s []int) {
	fmt.Println(len(s), cap(s))
}
