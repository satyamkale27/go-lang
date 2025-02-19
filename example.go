package main

import "fmt"

func main() {
	// variable()
	// arrays()
	//slice()
	//mapss()
	//function()
	//sum(1, 2)
	nextInt := intSeq()    /// `nextInt` refers to the closure returned by `intSeq`
	fmt.Println(nextInt()) // Prints 1 (i was 0, now it's incremented to 1)

}
