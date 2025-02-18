package main

import (
	"fmt"
)

func mapss() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("maps", m)

	v1 := m["k1"] //Gets a value for a key with name[key].
	fmt.Println("v1", v1)

	v3 := m["k3"] //  If the key doesn’t exist, the zero value of the value type is returned.
	fmt.Println("v3", v3)

	fmt.Println("len:", len(m)) //  The builtin len returns the number of key/value pairs when called on a map

	delete(m, "k1") //  The builtin delete removes key/value pairs from a map.
	fmt.Println("delete", m)

	clear(m) // To remove all key/value pairs from a map, use the clear builtin.
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
}
