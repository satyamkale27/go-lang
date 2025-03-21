package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	//  (int, error) specifies that the function returns two values: an int and an error.
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func main() {
	result, err := f(42)
	if err != nil {
		fmt.Println("Error:", err)

	} else {
		fmt.Println("Result:", result)
	}
}
