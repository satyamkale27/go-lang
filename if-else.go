package main

import "fmt"

func conditions() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 os odd")
	}

	if num := 10; num < 0 { // declared the num and same time compared
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println("it has 1 digit")
	} else {
		fmt.Println("it has multiple digits")
	}

}
