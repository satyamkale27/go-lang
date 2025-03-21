package main

import "fmt"

func rangexample() {
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums { // range iterates over elements in a variety of built-in data structures
		fmt.Println(num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for i, c := range "hello" {
		fmt.Println(i, c) // range on strings iterates over Unicode code points
	}

}

func main() {
	rangexample()
}
