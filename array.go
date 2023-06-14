package main

import (
	"fmt"
)

func main() {
	// var array [3]int8
	// fmt.Println(array[0], reflect.TypeOf(array[0]))

	// array[0] = 10
	// fmt.Println(array[0], reflect.TypeOf(array[0]))

	// array[0] = 11
	// fmt.Println(array[0], reflect.TypeOf(array[0]))
	// fmt.Println()

	// array[1] = 22
	// array[2] = 33
	// fmt.Println(array, reflect.TypeOf(array))

	// names := [3]string{"Adlet", "Madi", "Kayrat"}
	// fmt.Println(names, reflect.TypeOf(names))

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	marks := [5]int{5, 3, 3, 4, 2}
	sum := 0

	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}

	fmt.Println(float64(sum) / float64(len(marks)))
}
