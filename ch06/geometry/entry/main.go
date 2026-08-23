package main

import (
	"fmt"

	"github.com/yuhaowin/go-learning/ch06/geometry"
)

func main() {

	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	fmt.Println(geometry.Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))           // "5", method call

	perimeter := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perimeter.Distance()) // 周长为 "12"
}
