package main

import (
	"fmt"
)

func slice() {
	var s []string // declared nil slice
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)                                  // initialized slice
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // cap is capicity

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s) // gets all elements in slice
}
