package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Hello world " + fmt.Sprint(i))
	// }

	// var i int8
	// for i = 1; i <= 5; {
	// 	fmt.Print("Madi " + fmt.Sprint(i) + "\n")
	// 	i++
	// }

	// for i := 0; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// var num int16
	// for num = 0; num <= 10; num++ {
	// 	if num%2 != 0 {
	// 		continue
	// 	}

	// 	fmt.Println(num)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 		// continue
	// 	}

	// 	fmt.Println(i)
	// }

	// var name string
	// for true {
	// 	fmt.Scan(&name)
	// 	if name == "Adlet" {
	// 		break
	// 	}

	// 	fmt.Println("Hello " + name)
	// }

	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(nums, reflect.TypeOf(nums))
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// nums := []int{10, 20, 30, 40, 50}
	// for index, value := range nums {
	// 	fmt.Printf("Index - %d, Value - %d\n", index, value)
	// }

	// nums := []int{10, 20, 30, 40, 50}
	// for _, value := range nums {
	// 	fmt.Printf("Value - %d\n", value)
	// }

	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		fmt.Printf("%d ", matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, i := range matrix {
		for _, j := range i {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
