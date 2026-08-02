package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/yuhaowin/go-learning/flowup/functional/fibonacci"
)

func tryDefer1() {
	// defer 是被放到 stack 中，先进后出，执行的时候先打印 2 在打印 1
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
}

func tryDefer2() {
	for i := range 10 {
		defer fmt.Println(i) // 参数在 defer 语句时计算
		if i == 5 {
			panic("printed too many")
		}
	}
}

func writeFile1(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fib := fibonacci.Fibonacci()
	for range 20 {
		fmt.Fprintln(writer, fib())
	}
}

func writeFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")

	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s", pathError.Op, pathError.Path, pathError.Err)
		} else {
			panic(err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fib := fibonacci.Fibonacci()
	for range 20 {
		fmt.Fprintln(writer, fib())
	}
}

func main() {
	//tryDefer1()
	//tryDefer2()
	//writeFile1("123.txt")
	writeFile2("123.txt")
}
