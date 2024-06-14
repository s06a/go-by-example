package main

// closures are functions that can access or manipulate values
// from the outer function. also, here we use anonymous function,
// which is useful when you want to define a function without a name.

import "fmt"

// intSeq returns a closure that generates sequential integers.
// The closure keeps track of the last integer and increments it.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// factorial returns a closure that calculates the factorial of a given integer.
// The inner function maintains the result across multiple calls.
func factorial() func(int) int {
	var result int
	return func(num int) int {
		if num <= 1 {
			return 1
		}
		result = 1
		for i := 2; i <= num; i++ {
			result *= i
		}
		return result
	}
}

// multiplyBy returns a closure that multiplies a given factor with a number.
// The returned function accepts an integer and returns the product.
func multiplyBy(factor int) func(int) int {
	return func(num int) int {
		return factor * num
	}
}

func main() {
	// Create a closure for generating sequential integers
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())

	// Create a closure for calculating factorial
	returnFactorial := factorial()

	fmt.Println(returnFactorial(5))
	fmt.Println(returnFactorial(10))

	// Create a closure for multiplying by two
	multiplyByTwo := multiplyBy(2)
	fmt.Println(multiplyByTwo(5))
}
