package main

import "fmt"

func main() {
	// var name string
	// fmt.Scan(&name)

	// switch name {
	// case "Adlet":
	// 	fmt.Println("Hello adult son!")
	// case "Madi":
	// 	fmt.Println("Hello middle son!")
	// case "Kayrat":
	// 	fmt.Println("Hello little son!")
	// }

	// for true {
	// 	var name string
	// 	fmt.Scan(&name)

	// 	switch name {
	// 	case "Adlet":
	// 		fmt.Println("Hello adult son!")
	// 	case "Madi":
	// 		fmt.Println("Hello middle son!")
	// 	case "Kayrat":
	// 		fmt.Println("Hello little son!")
	// 	default:
	// 		fmt.Println("You are not my son!")
	// 	}
	// }

	// если один кейс сработает то остальные не сработает если даже будет тоже истинным, от этого можно избежать с помощью fallthrough
	var num int8
	fmt.Scan(&num)
	switch {
	case num > 1:
		fmt.Println("NUM greater than 1")
	case num < 11:
		fmt.Println("NUM less than 11")
	}
}
