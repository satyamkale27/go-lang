package main

import (
	"fmt"
	"unicode/utf8"
)

func stringg() {

	const s = "สวัสดี"
	fmt.Println("len(s):", len(s))

	for i := 0; i < len(s); i++ { // Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
		fmt.Println(s[i])
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%d. %c\n", idx, runeValue) // A range loop handles strings specially and decodes each rune along with its offset in the string.
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d. %c\n", i, runeValue)
		w = width // A string in Go is just a sequence of bytes.
		fmt.Println("width:", width)
		examineRunes(runeValue)
	}

}

func examineRunes(r rune) {
	if r == 't' {
		fmt.Println("found tee")

	} else if r == 'ส' {
		fmt.Println("found sua")
	}
}
