package main

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Distance 这里是值接收者方法，Distance 方法不会修改 Point 的值。
// 这里的 p Point 是一个值接收者，方法调用时会复制 Point 的值。
// Go 要求：
// 1. 方法的接收者类型必须是定义在同一个包里的类型，receiver 的类型必须是 T 或 *T
// 2. 方法的接收者类型不能是接口类型。
// 3. 方法的接收者类型不能是指针类型。
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

// ScaleBy 这里是指针接收者方法，ScaleBy 方法会修改 Point 的值。
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	// -----------------------------------------------------------------------------------------------------------------

	methodExpression()

	const day = 24 * time.Hour
	fmt.Println(day.Seconds()) // "86400"

}

func methodExpression() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // method expression
	fmt.Println(p.Distance(q))   // "5"
	fmt.Println(distance(p, q))  // "5" 这种函数会将其第一个参数用作接收器
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy // method expression
	scale(&p, 2)              // 这种函数会将其第一个参数用作接收器
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}
