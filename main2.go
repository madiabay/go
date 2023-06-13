package main

import "fmt"

func main() {
	// var num int64
	// fmt.Scan(&num)
	// if num > 0 {
	// 	fmt.Println("Num > 0")
	// } else if num < 0 {
	// 	fmt.Println("Num < 0")
	// } else {
	// 	fmt.Println("Num = 0")
	// }

	// if num >= 0 {
	// 	fmt.Println("Good")
	// } else {
	// 	fmt.Println("Bad")
	// }

	var name string
	fmt.Scan(&name)
	if len(name) > 10 {
		fmt.Println("Bad name")
	} else if len(name) < 10 {
		fmt.Println("Good name")
	} else {
		fmt.Println("OK")
	}
}
