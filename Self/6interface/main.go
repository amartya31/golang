/* Interfaces are named collections of method signatures. */

package main

import (
	"fmt"
	"math"
)

type Rec struct {
	Ht, Wd float64
}

func (r Rec) Area() float64 {
	return r.Ht * r.Wd
}
func (r Rec) Peri() float64 {
	return 2 * (r.Ht + r.Wd)
}

type Cir struct {
	rad float64
}

func (c Cir) Area() float64 {
	return math.Pi * math.Pow(c.rad, 2)
}
func (c Cir) Peri() float64 {
	return 2 * math.Pi * c.rad
}

/* CREATING INTERFACE  */
type Geometry interface {
	Area() float64 // -> note the return type has to be mentioned and the types should be same
	Peri() float64 // -> same here
}

func Measure(g Geometry) {
	fmt.Printf("The type of data [%T] with value [%+v]\n", g, g)
	area := g.Area()
	fmt.Printf("The Area calculated      = [%v]\n", area)
	peri := g.Peri()
	fmt.Printf("The Perimeter calculated = [%v]\n", peri)
}

func main() {
	r := Rec{2, 4}
	c := Cir{1}
	Measure(r)
	Measure(c)
	fmt.Println("\n\n")
}
