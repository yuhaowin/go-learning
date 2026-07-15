package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// slice 是 arr 的视图

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	fmt.Println("after update")
	update(s1)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arr2[:]
	fmt.Println(s3)
	fmt.Println("after reslice")
	s3 = s3[2:]
	fmt.Println(s3)
	s3 = s3[:5]
	fmt.Println(s3)

	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(len(arr3))
	fmt.Println(cap(arr3))

	//s4的值为［2 3 4 5］，s5的值为［5 6］
	//slice可以向后扩展，不可以向前扩展
	//s[i] 不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
	s4 := arr3[2:6] // 2 3 4 5
	s5 := s4[3:5]
	fmt.Println("Extending slice")
	fmt.Println(s4)
	fmt.Println(s5)

	// slice 添加数据
	// 添加元素时如果超越cap，系统会重新分配更大的底层数组
	s6 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s7 := append(s6, 8)
	s8 := append(s7, 9)
	s9 := append(s8, 10)

	fmt.Println(s6, len(s6), cap(s6))
	fmt.Println(s7, len(s7), cap(s7))
	fmt.Println(s8, len(s8), cap(s9))
	fmt.Println(s9, len(s9), cap(s9))
}

func update(s []int) {
	s[0] = 100
}
