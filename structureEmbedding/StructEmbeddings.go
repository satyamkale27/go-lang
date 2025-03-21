package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
	// The base struct has a method describe with a receiver of type base.
	// This means that any instance of base can call the describe method.
}

type container struct {
	base
	str string

	// container struct embeds base
	// When the container struct embeds the base struct, it inherits the describe method from base. This allows an instance of container to call the describe method as if it were its own method.
}

func main() {

	co := container{ // An instance of container named co is created.
		base: base{
			num: 1,
		},
		str: "name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe()) // The container struct inherits the describe() method from the base struct

	type describer interface {
		describe() string
	}

	var d describer = co
	//  co assigns the variable co (which is an instance of the container struct) to the variable d of type describer.
	// Interface Satisfaction: The container struct inherits the describe() method from the base struct, so it satisfies the describer interface.
	fmt.Println("describer:", d.describe())

}
