package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	total := sum(1, 2, 3)
	fmt.Println(total)

	nums := []int{1, 2, 3, 4}
	// no need for ":" because we have declared it before
	total = sum(nums...)
	fmt.Println(total)
}
