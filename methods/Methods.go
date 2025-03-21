package main

import "fmt"

// A method in Go is just a function with a receiver. The receiver is what the method operates on
type rect struct {
	width, height int
}

func (r *rect) area() int {
	// The receiver (r *rect) is a pointer (*rect), meaning: It allows modifying the actual struct inside the method.
	//It avoids copying large structs, making it more efficient.
	// r.width = 16 // it modify original struct
	return r.width * r.height

}

func (r rect) perim() int {
	return 2*r.width + 2*r.height // Now r is a copy of the original struct. Any modifications inside area() won’t affect the original struct.
	// any modification done to width or height will not modify original instead modify the copy ones
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(r)
	fmt.Println(r.perim())

	rp := &r // Go automatically handles conversion between values and pointers for method calls.
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
