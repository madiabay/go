package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for i, el := range slice {
		fmt.Printf("%d: %d\n", i, el)
	}
}
