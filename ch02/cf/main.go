package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yuhaowin/go-learning/ch02/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fahrenheit := tempconv.Fahrenheit(t)
		celsius := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			fahrenheit, tempconv.FToC(fahrenheit),
			celsius, tempconv.CToF(celsius))
	}
}
