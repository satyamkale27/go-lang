package main

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
