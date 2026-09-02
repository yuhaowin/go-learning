package main

import (
	"bufio"
	"fmt"
	"os"
)

// Go 初始化空 Map
// make(map[K]V)
// m := make(map[string]int)
// go run main.go < sample.txt
// < sample.txt 把内容送给 stdin

// bufio 包，处理输入和输出
func main() {
	counts := make(map[string]int)
	fmt.Println(counts)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 每次调用 input.Scan()，即读入下一行，并移除行末的换行符，Scan 函数在读到一行时返回 true，不再有输入时返回 false。
		// 读取的内容可以调用 input.Text() 得到。
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count:%d\tline:%s\n", n, line)
		}
	}
}
