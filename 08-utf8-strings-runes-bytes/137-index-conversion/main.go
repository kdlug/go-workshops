package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	str := "Yūgen, Good morning. おはようございます。🤣"

	// a string is read only slice
	// str[0] = 'M' cannot be changed
	// str[1] = 'O'

	// we can covert it to a byte slice
	bytes := []byte(str)
	bytes[0] = 'M'
	bytes[1] = 'O'

	newStr := string(bytes)                                  // creates a string and copies the byte[] to the new strings backing array
	fmt.Printf("str (len: %d): %s\n", len(str), str)         // str: Yūgen, Good morning. おはようございます。 because bytes is a copy!
	fmt.Printf("newStr (len: %d): %s\n", len(bytes), newStr) // newStr: MO�gen , Good morning. おはようございます。

	// count bytes
	// byte lengths are equal
	fmt.Printf("str len: %d\n", len(str))     // str len: 56
	fmt.Printf("bytes len: %d\n", len(bytes)) // bytes len: 56

	// count runes -> 32/33 characters
	// so so runes have multiple bytes
	fmt.Printf("str len runes: %d\n", utf8.RuneCountInString(str)) // str len runes: 32
	fmt.Printf("bytes len runes: %d\n", utf8.RuneCount(bytes))     // bytes len runes: 33

	for i, r := range str { // jumps over the runes instead of the bytes in a string
		fmt.Printf("str[%2d] = %-12x = %q\n", i, string(r), r)
	}
	// OUTPUT:
	// str[ 0] = 59
	// str[ 1] = c5
	// str[ 3] = 67
	// ...
	// -> some indices are missing, because some characters are stored in 2 bytes

	// indexes
	fmt.Printf("1st byte: %c\n", str[0])
	fmt.Printf("2nd byte: %c\n", str[1])   // 2nd byte: Å -> because 2nd character has 2 bytes
	fmt.Printf("2nd rune: %s\n", str[1:3]) // 2nd char: ū // slicing

	// RUNES
	// we can always convert a string to a rune slice
	// which is less efficient
	runes := []rune(str)
	fmt.Printf("runes: %s\n", string(runes))
	fmt.Printf("runes len runes: %d\n", len(runes))                              // 32
	fmt.Printf("runes bytes len: %d\n", int(unsafe.Sizeof(runes[0]))*len(runes)) // 128 bytes -> string has 56 bytes

	fmt.Printf("1st rune: %c\n", runes[0])
	fmt.Printf("2nd rune: %c\n", runes[1]) // 2nd byte: Å -> because 2nd character has 2 bytes
	fmt.Printf("2nd rune: %c\n", runes[2])
	fmt.Printf("First 5 runes: %c\n", runes[0:5])
}
