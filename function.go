package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multiReturnValues() (int, int) {
	return 1, 2
}

func vardicValues() (int, int) {
	return 5, 7
}

func function() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3) // res =   // note res :=  it declared and initialise ,  res = it reassign the value, now res will overwrite the value of previous res with new one
	fmt.Println("1+2+3 =", res)
	a, b := multiReturnValues() // store multiple value from function
	fmt.Println("res =", a, b)
	_, c := vardicValues() // If you only want a subset of the returned values, use the blank identifier _.
	fmt.Println("res =", c)
}
