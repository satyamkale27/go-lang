package main

import "fmt"

func variable() {
	// var can declare one or more variables at ones //
	var a = "initial"
	fmt.Println(a)
	var b, c int = 3, 4
	fmt.Println(b, c)
	var s = "satyam" // Go, variables can be declared with or without specifying the data type //
	fmt.Println(s)

	f := "shorthand"
	fmt.Println(f) // it is shorthand for declaring and initializing eg:- var f = "shorthand" //

	const pi = 3.14 // constant are immutable, they cannot be changed //
	fmt.Println(pi)
}

func main() {

	variable()

}
