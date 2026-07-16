package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep *string = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(n, sep)
	fmt.Println(*n, *sep)
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
