package main

import "fmt"

// define a new type 'rectangle'
type rectangle struct {
	width, height int
}

// define a method 'Area' for the 'rectangle' type
func (r rectangle) Area() int {
	return r.width * r.height
}

// define a new type 'counter'
type counter struct {
	value int
}

// define a method 'Increment' with a pointer receiver
func (c *counter) Increment() {
	c.value++
}

func (c *counter) Decrement() {
	c.value--
}

type myInt int

func (m myInt) IsEven() bool {
	return m%2 == 0
}

func main() {
	r := rectangle{width: 10, height: 5}
	fmt.Println("area:", r.Area())

	c := counter{value: 0}
	c.Increment()
	fmt.Println("value of c is", c.value)

	num := myInt(4)
	fmt.Println(num, "is even?", num.IsEven())
}
