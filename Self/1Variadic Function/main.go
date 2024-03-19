/*Variadic functions can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function. */

package main

import (
	"fmt"
)

func main() {
	num := []uint{1, 2, 3, 4}
	var total uint = Add(num...)
	fmt.Printf("Total of [%v] = [%v]\n", num, total)
	total = Add(5, 6, 7, 8)
	fmt.Printf("Total of [%v] = [%v]\n", "5, 6, 7, 8", total)

}
func Add(num ...uint) uint {
	var total uint = 0
	for _, val := range num {
		total += val
	}
	fmt.Printf("The type of num = [%T]\n ", num)
	return total
}
