/* Points to find -
1) Embedding syntax -
2) Field inheritance -
3) Methord inheritance
4) Name collision -> filed and methord overiding
5) Accessing
6) Closures  & Anonymous function
*/

package main

import "fmt"

type Animal struct { // -> like base class
	Name string
	Age  uint
	gen  bool
}

func (a Animal) PrintGeneric() {
	fmt.Printf("The Data for the value is [%+v]\n", a)
}

// Structure Embedding
type Dog struct {
	Animal // -> When you omit the field name when embedding, the type name becomes the field name.
	Breed  string
}
type Cat struct {
	Charac Animal // -> When you keep the name, that name is used
	Breed  string
}

// Printing The new struct
func (d Dog) PrintAll() {
	fmt.Printf("The whole structure Dog is [%+v]\n", d)
}
func (c Cat) PrintAll() {
	fmt.Printf("The whole structure Cat is [%+v]\n", c)
}

// Interface
type PrintGenericI interface {
	PrintAll()
}

func PrintData(g PrintGenericI) {
	g.PrintAll()
}

// closures
func AddAll() func() uint {
	var i uint = 0
	return func() uint {
		i++
		return i
	}

}
func main() {
	d := Dog{Animal: Animal{"Bud", 5, true},
		Breed: "Doverman"}
	fmt.Printf("The entry for Dog is type [%T] and data [%+v]\n", d, d)
	c := Cat{Charac: Animal{"Cute", 2, false},
		Breed: "Wild"}
	CallAdd := AddAll()
	TotalAnimals := 0
	fmt.Printf("The entry for Cat is type [%T] and data [%+v]\n", c, c)
	fmt.Println("Methord Inheritance of Dog")
	d.PrintGeneric() // -> NOTE DIRECTLY CALLED
	fmt.Println("Methord Inheritance of Cat")
	c.Charac.PrintGeneric() // ->  INDIRECTLY CALLED
	fmt.Println("Interface Dog")
	PrintData(d)
	TotalAnimals = int(CallAdd())
	fmt.Println("Interface Cat")
	PrintData(c)
	TotalAnimals = int(CallAdd())
	fmt.Printf("The total animals are [%v]\n", TotalAnimals)
}
