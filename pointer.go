package main

import (
	"fmt"
)

// func main() {
// 	a := 5
// 	// pointer := &a
// 	fmt.Println(*&a)
// }

// func change(str *string) {
// 	*str = "Madi"
// }

// func main() {
// 	s := "Abay"
// 	fmt.Println(s)
// 	change(&s)
// 	fmt.Println(s)
// }

// func main() {
// 	name := "Madi"
// 	fmt.Println(name)
// 	*&name = "Abay"
// 	fmt.Println(name)
// }

func main() {
	a := 10
	b := &a
	fmt.Println(a, b)
	*b = 20
	fmt.Println(a, b)
}
