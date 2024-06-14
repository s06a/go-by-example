package main

import (
	"fmt"
)

// maps are like hashes and dicts in other languages

func main() {
	m := make(map[string]int) // string keys and integer values

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println(m, len(m))

	// delete by key
	delete(m, "k2")

	// delete all
	clear(m)

	// returns value and an indicator if key is present or not
	_, prs = m["k3"]
	fmt.Println(prs)

	// like slices here we also have maps.Equal(m1, m2)
}
