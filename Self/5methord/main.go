package main

import "fmt"

type Shape struct {
	Height uint
	Width  uint
}

func (s Shape) Area() (int, string) {
	shape := "Square"
	if s.Height != s.Width {
		shape = "Rectangle"
	}
	var area int = int(s.Height * s.Width)
	return area, shape
}

func (a *Shape) MakeSquare() {
	if a.Height > a.Width {
		a.Height = a.Width
	} else {
		a.Width = a.Height
	}
}

func main() {
	a := Shape{12, 3}
	area, shape := a.Area()
	fmt.Printf("Area of [%v] is [%v]\n", shape, area)
	a.MakeSquare()
	area, shape = a.Area()
	fmt.Printf("Area of [%v] is [%v]\n", shape, area)
}
