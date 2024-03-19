/* Interface of a function add which returns int in one case and string in case of another
and print separately */

package main

import (
	"fmt"
)

type IntVal struct {
	a, b int
}

func (i IntVal) Add() interface{} { // RETURN TYPE HAS TO BE INTERFACE
	res := i.a + i.b
	fmt.Printf("Adding the values [%+v] with res [%v]\n", i, res)
	return res
}

type StrVal struct {
	a, b string
}

func (s StrVal) Add() interface{} { // RETURN TYPE HAS TO BE INTERFACE
	res := s.a + s.b
	fmt.Printf("Adding the values [%+v] with res [%v]\n", s, res)
	return res
}

/* CREATING INTERFACE */
type AddItem interface {
	Add() interface{}
}

func PrintValue(a AddItem) {
	result := a.Add()

	switch result.(type) { //  TYPE CHECK THE RETURN TYPE
	case int:
		fmt.Printf("The type is INT \n")
		fmt.Printf("The result is [%v]\n", result)
	case string:
		fmt.Printf("The type is STRING \n")
		fmt.Printf("The result is [%v]\n", result)
	}
}

func main() {
	i := IntVal{1, 2}
	s := StrVal{"A", "B"}
	PrintValue(i)
	PrintValue(s)
}
