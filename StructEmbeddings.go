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
