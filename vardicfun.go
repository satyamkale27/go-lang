package main

import "fmt"

func sum(nums ...int) { // this function can take arbitrary number of ints as argument, a slice can be passed as argument
	fmt.Println(nums, "v")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
