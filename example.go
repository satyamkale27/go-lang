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

	//var fib func(n int) int //  Anonymous functions can also be recursive,
	//// but this requires explicitly declaring a variable with var to store the function before it’s defined.
	//fib = func(n int) int {
	//	if n < 2 {
	//		return n
	//	}
	//	return fib(n-1) + fib(n-2)
	//}
	//fmt.Println(fib(7))

	// rangexample()

	//i := 1
	//fmt.Println("initial", i)
	//zeroval(i)
	//fmt.Println("zeroval", i)
	//zeroptr(&i)
	//fmt.Println("zeroptr", i)

	//stringg()

	//fmt.Println(person{"Bob", 20})
	//fmt.Println(person{name: "Bob", age: 42})
	//fmt.Println(person{name: "Fred"})
	//fmt.Println(&person{name: "Ann", age: 40})
	//fmt.Println(newPerson("Jon"))
	//
	//s := person{name: "Sean", age: 50}
	//fmt.Println(s.name) // Access struct fields with a dot.
	//s.age = 19          // mutates an instance
	//fmt.Println(s.age)  // structs are mutable
	//
	//dog := struct {
	//	name   string
	//	isGood bool
	//}{
	//	"Rex",
	//	true,
	//}
	//fmt.Println(dog)

	r := rect{width: 10, height: 5}
	fmt.Println(r)
	fmt.Println(r.perim())

	rp := &r // Go automatically handles conversion between values and pointers for method calls.
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

}
