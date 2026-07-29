package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes中国北京!"
	fmt.Println(len(s))
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}

	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %x) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	// r 是 bytes 转出的第一个 rune
	// size 是这个 rune 占用的字节数量
	r, size := utf8.DecodeRune(bytes)
	fmt.Printf("%c %d\n", r, size)

	for len(bytes) > 0 {
		r, size = utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", r)
	}

	fmt.Println()

	// rune 相当于是 go 的 char

	for _, r := range []rune(s) {
		fmt.Printf("%c ", r)
	}
}
