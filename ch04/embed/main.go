package main

import "fmt"

type Point struct{ X, Y int }

// Circle 只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。
type Circle struct {
	Point  //匿名成员
	Radius int
}

type Wheel struct {
	Circle // 匿名成员
	Spokes int
}

func main() {
	var w0 Wheel
	w0.X = 8      // equivalent to w.Circle.Point.X = 8
	w0.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w0.Radius = 5 // equivalent to w.Circle.Radius = 5
	w0.Spokes = 20

	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
}
