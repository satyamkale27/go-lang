package main

import "fmt"

func main() {
	// variable()
	// arrays()
	//slice()
	//mapss()
	//function()
	//sum(1, 2)
	//nextInt := intSeq()    /// `nextInt` refers to the closure returned by `intSeq`
	//fmt.Println(nextInt()) // Prints 1 (i was 0, now it's incremented to 1)

	// fmt.Println(fact(3))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
