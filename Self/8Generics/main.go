package main

import "fmt"

type Num interface {
	int | int16 | int64 | float32 | float64
}

func Add[T Num](a T, b T) T {
	return a + b
}

func main() {
	fmt.Printf("Trying Generics ")
	fmt.Printf("The add of int   1   + 2   = [%v]\n", Add(1, 2))
	fmt.Printf("The add of float 1.1 + 2.2 = [%v]\n", Add(1.1, 2.2))

}
