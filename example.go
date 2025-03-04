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

	//r := rect{width: 10, height: 5}
	//fmt.Println(r)
	//fmt.Println(r.perim())
	//
	//rp := &r // Go automatically handles conversion between values and pointers for method calls.
	//fmt.Println(rp.area())
	//fmt.Println(rp.perim())

	//r := recta{width: 3, height: 4}
	//c := circle{radius: 5}

	//measure(r) // Calls measure() with a rectangle
	//measure(c) // Calls measure() with a circle
	//
	//detectCircle(r)
	//detectCircle(c)

	//co := container{ // An instance of container named co is created.
	//	base: base{
	//		num: 1,
	//	},
	//	str: "name",
	//}
	//
	//fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	//fmt.Println("also num:", co.base.num)
	//fmt.Println("describe:", co.describe()) // The container struct inherits the describe() method from the base struct
	//
	//type describer interface {
	//	describe() string
	//}
	//
	//var d describer = co
	////  co assigns the variable co (which is an instance of the container struct) to the variable d of type describer.
	//// Interface Satisfaction: The container struct inherits the describe() method from the base struct, so it satisfies the describer interface.
	//fmt.Println("describer:", d.describe())

	//ints := []int{1, 2, 3, 4}
	//strs := []string{"a", "b", "c", "d"}
	//
	//fmt.Println(SlicesIndex(ints, 3))
	//fmt.Println(SlicesIndex(strs, "c"))
	//
	//lst := List[int]{}
	//lst.Push(10)
	//lst.Push(13)
	//lst.Push(23)
	//
	//fmt.Println("list", lst.AllElements())

	result, err := f(42)
	if err != nil {
		fmt.Println("Error:", err)

	} else {
		fmt.Println("Result:", result)
	}

}
