package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type Numbers_Interface interface {
	Sum() int
	Mul() int
	Div() float64
	Sub() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Mul() int {
	return n.num1 * n.num2
}

func (n Numbers) Div() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n Numbers) Sub() int {
	return n.num1 - n.num2
}

func main() {
	var i Numbers_Interface
	numbers := Numbers{num1: 15, num2: 10}
	i = numbers
	fmt.Printf("Sum: %d\n", i.Sum())
	fmt.Printf("Mul: %d\n", i.Mul())
	fmt.Printf("Div: %f\n", i.Div())
	fmt.Printf("Sub: %d\n", i.Sub())
}
