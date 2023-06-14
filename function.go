package main

import "fmt"

// func first_func() {
// 	fmt.Println("Hello World")
// }
// func main() {
// 	first_func()
// 	first_func()
// 	first_func()
// 	first_func()
// }

// func sum(a int, b int) {
// 	fmt.Println(a + b)
// }
// func main() {
// 	sum(19, 20)
// }

// func square_num(num int) int {
// 	return int(math.Pow(float64(num), 2))
// }
// func main() {
// 	fmt.Println(square_num(12))
// }

// func plus_minus(num1 int, num2 int) (int, int) {
// 	plus := num1 + num2
// 	minus := num1 - num2
// 	return plus, minus
// }
// func main() {
// 	p, m := plus_minus(15, 10)
// 	fmt.Printf("Plus result-%d, Minus result-%d\n", p, m)
// }

func tim_div(a int, b int) (times int, div int) {
	times = a * b
	div = a / b
	return
}
func main() {
	a, b := tim_div(15, 5)
	fmt.Println(a, b)
}
