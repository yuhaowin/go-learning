package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] // os.Args[0] 是命令名称，os.Args[1:] 表示 1 - len 的 slice
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// map 和 struct 的值传递区别：
//   - struct（如 os.File）是值类型，传值会复制整个结构体，
//     函数内部拿到的是独立副本，修改不影响外部，所以 f 必须传指针 *os.File。
//   - map 是引用类型，变量内部是指向底层哈希表的描述符（含指针），
//     传值时只复制描述符，底层数据仍是同一份，所以 counts 不用传指针，
//     函数内部对它的增删改（如 counts[key]++）外部也能看到。
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
