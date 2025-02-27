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

/*


UTF-8 Encoding of "สวัสดี"
Here is how each character gets encoded in UTF-8:

ส (U+0E2A):

Encodes to 0xE0 0xB8 0xAA (3 bytes).
ว (U+0E27):

Encodes to 0xE0 0xB8 0xA7 (3 bytes).
ั (U+0E35):

Encodes to 0xE0 0xB8 0xA5 (3 bytes).
ส (U+0E2A):

Repeated, encodes to 0xE0 0xB8 0xAA (3 bytes).
ด (U+0E34):

Encodes to 0xE0 0xB8 0x94 (3 bytes).
ี (U+0E35):

Repeated, encodes to 0xE0 0xB8 0xA5 (3 bytes).


UTF-8 Byte Representation:
The entire string "สวัสดี" is encoded in the following byte sequence:


0xE0 0xB8 0xAA 0xE0 0xB8 0xA7 0xE0 0xB8 0xA5 0xE0 0xB8 0xAA 0xE0 0xB8 0x94 0xE0 0xB8 0xA5
Each of these groups of three bytes represents a single Thai character.



How utf8.RuneCountInString(s) Works:
utf8.RuneCountInString(s) counts how many Unicode code points (runes) there are in the string, and it decodes each rune from the byte sequence based on the leading bits of the bytes.

First Byte Analysis:

The first byte of the string is 0xE0, which starts with 1110. This means it is the start of a 3-byte character.
The function identifies that the next two bytes belong to the same rune, and together they form the character "ส" (U+0E2A).
The function counts this as 1 rune.
Second Byte Analysis:

The next byte is 0xE0 (again), indicating the start of another 3-byte character.
The function decodes the next two bytes to get the character "ว" (U+0E27), counting it as 1 rune.
Third Byte Analysis:

The next byte is 0xE0 again, and the function decodes the next two bytes to get the character "ั" (U+0E35), counting it as 1 rune.
Repeat for Remaining Characters:

The function continues this process for each subsequent character in the string, identifying the start of each rune and decoding the multi-byte sequence.
By the end, the function processes all 6 characters in the string and counts 6 runes.

Total Rune Count:
utf8.RuneCountInString("สวัสดี") will return 6 because there are 6 Thai characters in the string, each of which is a separate rune.



*/
