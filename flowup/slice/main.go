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
}

func update(s []int) {
	s[0] = 100
}
