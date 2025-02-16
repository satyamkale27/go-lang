package main

import (
	"fmt"
)

func slice() {
	var s []string // declared nil slice
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)                                  // initialized slice  // Slice created with `make`
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // cap is capicity

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s) // gets all elements in slice

	s = append(s, "d") // add more element to slice
	s = append(s, "e", "f")
	fmt.Println("append", s)

	/*  *** working ***
		 - The slice `s` has length 3 and capacity 3 at this point. Adding one more element exceeds its capacity.

		    1. A new underlying array is allocated with double the previous capacity (by default in Go).
		    2. All existing elements ("a", "b", "c") are copied to the new array.
		    3. The new element "d" is added.

	       - The slice now becomes:
		    - Content: `s = ["a", "b", "c", "d"]`
		    - Length: 4
		    - Capacity: 6 (doubled from 3, but this depends on Go's internal resizing logic).

		  s = append(s, "e", "f")
		  - Now, two more elements "e" and "f" are appended to the end of the slice.
		  - Since the current capacity (6) can accommodate these elements without reallocation:

		    1. The slice adds the new elements "e" and "f" in the available space of the existing underlying array.


		 - The slice now becomes:
		    - Content: `s = ["a", "b", "c", "d", "e", "f"]`
		    - Length: 6
		    - Capacity: 6 (unchanged since there's enough space).

	*/

	c := make([]string, len(s)) // it creates the same slice as of s
	copy(c, s)
	fmt.Println("copies c from s", c)

	l := s[2:5] //  l := s[2:5]` creates a new slice containing the elements `s[2]`, `s[3]`, and `s[4]`.
	fmt.Println("sl1", l)

	t := []string{"g", "h", "i"} // We can declare and initialize a variable for slice in a single line as well.
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {

		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // it will create a nill slice at i th position

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j // allocate the slices that were created
		}

	}

	fmt.Println("twoD slice", twoD)

}
