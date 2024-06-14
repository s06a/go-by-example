package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a) // default value is array of zeros

	a[4] = 10
	fmt.Println(a)
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// auto count
	c := [...]int{1, 2, 3}
	fmt.Println(c)

	d := [...]int{2: 1, 1: 2}
	fmt.Println(d)

	var twoD [4][4]int
	for i := range 4 {
		for j := range 4 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
	fmt.Println(twoD[2][:]) // third row
}
