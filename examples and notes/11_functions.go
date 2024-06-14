package main

import (
	"errors"
	"fmt"
)

func calculator(operation string, a, b int) (int, error) {
	if operation == "add" {
		return a + b, nil
	}
	return 0, errors.New("Invalid operation")
}

func main() {
	result, err := calculator("add", 10, 4)
	if err == nil {
		fmt.Println("result is", result)
	} else {
		fmt.Println(err)
	}
}
