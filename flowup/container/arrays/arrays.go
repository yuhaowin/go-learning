package main

import "fmt"

func main() {

	var arr1 [5]int
	// 短变量声明，必须赋初值
	arr2 := [3]int{1, 3, 5}
	// ... 会自动数数组的长度
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(len(arr3))

	// 4 行 5 列，4 个长度为 5 的数组
	var grid [4][5]int
	fmt.Println(grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for index, value := range arr3 {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	printArr(arr1)
	// cannot use arr2 (variable of type [3]int) as [5]int value in argument to printArr
	//printArr(arr2)
	printArr(arr3)

	modifyArr(&arr3)
	fmt.Println(arr3) //[100 4 6 8 10]
}

func printArr(arr [5]int) {
	// 数组是值传递，在函数内部修改，不会反应到外部
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

func modifyArr(arr *[5]int) {
	// 正常来说 arr 是 *[5]int 类型（指向数组的指针），按 C 的写法你得先解引用再取下标，比如 (*arr)[0] = 100。但 Go 允许直接写 arr[0]，编译器会自动帮你解引用，等价于：
	// (*arr)[0] = 100
	(*arr)[0] = 100 // 显式解引用，合法
	arr[0] = 100    // 语法糖，等价写法
}
