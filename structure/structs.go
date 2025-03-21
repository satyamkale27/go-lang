package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person { // *person means Function returning pointer

	// newPerson(name string) *person → Creates a new person instance with age = 42 and returns a pointer to it.
	p := person{name: name}
	p.age = 42
	return &p // returns pointer
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob", age: 42})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Access struct fields with a dot.
	s.age = 19          // mutates an instance
	fmt.Println(s.age)  // structs are mutable
}
