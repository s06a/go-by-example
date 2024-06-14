package main

import "fmt"

func main() {

	var b int = 8

	if a := 7; a%2 == 0 && b%2 == 0 {
		fmt.Println("both are even")
	} else if a%2 == 0 || b%2 == 0 {
		fmt.Println("a or b is even")
	} else {
		fmt.Println("both are odd")
	}

}
