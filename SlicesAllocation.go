package main

import "fmt"

func main() {
	innerlength := 3
	var twoD [][]int
	twoD = make([][]int, 3)

	for i := 0; i < 3; i++ {
		// declare slices inside slice
		twoD[i] = make([]int, innerlength)
	}

	// allocate slices

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
