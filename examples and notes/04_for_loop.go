package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			break
		} else if i%2 == 0 {
			continue
		}
	}

	j := 0
	for j <= 3 {
		fmt.Println(j)
		j += 1
	}

}
