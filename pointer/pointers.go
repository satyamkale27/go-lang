package main

import "fmt"

func zeroval(ival int) { // zeroval has an int parameter, so arguments will be passed to it by value. zeroval will get a copy of ival
	ival = 0
}

func zeroptr(iptr *int) { // changes value of variable  // // Dereferencing the pointer and modifying the original value
	*iptr = 0 // modifying the value
}

func main() {
	i := 1
	fmt.Println("initial", i)
	zeroval(i)
	fmt.Println("zeroval", i)
	zeroptr(&i)
	fmt.Println("zeroptr", i)
}
