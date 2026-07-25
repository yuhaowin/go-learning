package main

import (
	"fmt"
	"golearning/flowup/functional/fibonacci"
)

func main() {
	f := fibonacci.Fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
