package main

// string is a read-only slice of bytes, In Go, the
// concept of a character is called a rune - it’s an
// integer that represents a Unicode code point.

import "fmt"

func main() {
	str := "hello world"

	fmt.Println("len of string:", len(str))

	for i, r := range str {
		fmt.Printf("index: %d, rune: %c\n", i, r)
		examineRune(r)
	}
}

func examineRune(r rune) {
	if r == 'o' {
		fmt.Println("found o")
	}
}
