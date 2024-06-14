package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	// use _ if you don't need index
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "bananan"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys:", k)
	}

	// iterate over unicode code points
	for _, c := range "string" {
		fmt.Println("word:", c)
	}
}
