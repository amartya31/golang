package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	a := fact(7)
	fmt.Printf("The value = [%v]\n", a)
	/* Closures can also be recursive, but this requires the closure to
	be declared with a typed var explicitly before it’s defined. */
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Printf("The value = [%v]\n", fib(7))
}
