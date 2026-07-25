package main

import "fmt"

// 1, 1, 2, 3, 5, 8, 13
//
//	a, b
//	   a, b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
