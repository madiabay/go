package main

import (
	"fmt"
)

func main() {
	// a := 10
	// if a >= 1 {
	// 	fmt.Println("a >= 1, " + "a=" + fmt.Sprint(a))
	// }

	// var a int8 = 1
	// var b int8 = 127
	// if a == 1 && b > a {
	// 	fmt.Println("TRUE!!!")
	// }
	// if a != 2 || b < a {
	// 	fmt.Println("ALSO TRUE!!!")
	// }

	// var c int8 = 10
	// var d int16 = 1999
	// if d != int16(c) {
	// 	fmt.Println("c NOT EQUAL TO 10")
	// }

	var name string
	fmt.Scan(&name)

	if name == "Madi" || name == "Mado" || name == "Abdi" {
		fmt.Println("Your name is MADI")
	}
}
