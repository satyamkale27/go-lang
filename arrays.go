package main

import "fmt"

func arrays() {
	var a [5]int
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get:", a[4])
	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Print("dcl", b)

	b = [...]int{100, 3: 400, 500} // redeclares array with same length  The elements at indices 1 and 2 are not explicitly set, so they default to 0.
	fmt.Println(b)

	var twod [2][3]int

	twod = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(twod)

}
