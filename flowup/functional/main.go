package main

import (
	"fmt"

	"github.com/yuhaowin/go-learning/flowup/functional/fibonacci"
)

func main() {
	f := fibonacci.Fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
