package main

//import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // This fact function calls itself until it reaches the base case of fact(0).
}

/*
Additional note:-
Think of recursion as stacking function calls like this:
fact(3)  // Waiting for fact(2)
    fact(2)  // Waiting for fact(1)
        fact(1)  // Waiting for fact(0)
            fact(0)  // Returns 1 (Base case)
        fact(1) returns 1 * 1 = 1
    fact(2) returns 2 * 1 = 2
fact(3) returns 3 * 2 = 6
*/
