package main

import "fmt"

type Info struct {
	Name string
	Age  uint
}

func main() {
	var a Info = Info{Name: "Amartya", Age: 5}
	fmt.Printf("Value of a [%+v]\n", a)

	// CREATE OWN STRUCTURE
	b := struct {
		Name string
		Bois int
	}{
		"Anish",
		24,
	}
	fmt.Printf("Value of b [%+v]\n", b)

}
