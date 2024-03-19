/* Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function
inline without having to name it. */

package main

import (
	"fmt"
)

/*
	func intSeq() func() int {: This line declares a function named intSeq

that returns another function. The returned function is specified to have
no parameters and return an integer.
*/
func AddSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	a := AddSeq()
	fmt.Printf("The value of I = [%v]\n", a())
	fmt.Printf("The value of I = [%v]\n", a())
	fmt.Printf("The value of I = [%v]\n", a())
	b := AddSeq()
	fmt.Printf("The value of I = [%v]\n", b())
}
