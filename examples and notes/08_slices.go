package main

import (
	"fmt"
	"slices"
)

func main() {

	// uninitialized slice with zero length
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	// empty slice with non-zero length
	s = make([]string, 3)
	s[2] = "0"
	s = append(s, "1", "2", "3", "4")
	fmt.Println("capacity:", cap(s), "len:", len(s), "set:", s, "third element:", s[2])

	// make a copy
	c := make([]string, len(s))
	copy(c, s)
	// check if slices are equal
	if slices.Equal(c, s) {
		fmt.Println("c and s are equal")
	}

	// slices of slice
	d := c[2:5] // also [:a], [b:] works
	fmt.Println(d)

	// an empty 3X3 slice with multiple length inner slices
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
