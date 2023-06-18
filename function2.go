package main

import "fmt"

// func main() {
// 	a := func(name string) string {
// 		return "Hello " + name
// 	}
// 	fmt.Println(a("Madi"))
// }

// func some_function(some_func func(int) int) {
// 	fmt.Println(some_func(25))
// }
// func main() {
// 	square := func(x int) int {
// 		return x * x
// 	}
// 	some_function(square)
// }

// func some_func(str string) func() {
// 	return func() {
// 		fmt.Println("Hello", str)
// 	}
// }
// func main() {
// 	a := some_func("Madi")
// 	a()
// }

func some_func(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	some_func("Hello")()
}
