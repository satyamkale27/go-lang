package main

import "fmt"

// Define the describer interface
type describer interface {
	describe() string
}

// Define the baseStruct struct with a describe method
type baseStruct struct {
	num int
}

func (b baseStruct) describe() string {
	return fmt.Sprintf("baseStruct with num=%v", b.num)
}

// Define the containerStruct struct that embeds baseStruct
type containerStruct struct {
	baseStruct
	str string
}

// Define another struct itemStruct with a describe method
type itemStruct struct {
	id   int
	name string
}

func (i itemStruct) describe() string {
	return fmt.Sprintf("itemStruct with id=%v and name=%v", i.id, i.name)
}

// Function that accepts any type that implements the describer interface
func printDescription(d describer) { // d describer(describer)  check for the interface
	fmt.Println("describer:", d.describe())
}

func main() {
	co := containerStruct{
		baseStruct: baseStruct{
			num: 1,
		},
		str: "name",
	}

	it := itemStruct{
		id:   2,
		name: "example",
	}

	// Direct method call
	fmt.Println("describe:", co.describe())

	// Using the interface
	var d describer = co
	printDescription(d)

	// Using the interface with another type
	printDescription(it)
}
