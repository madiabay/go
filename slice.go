package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice := []int8{3, 1, 2}
	// array := [3]int8{3, 1, 2}
	// fmt.Println(slice, reflect.TypeOf(slice))
	// fmt.Println(array, reflect.TypeOf(array))
	// slice = append(slice, 0)
	// fmt.Print(slice, array, "\n")
	// slice := []string{"a", "b", "c"}
	// fmt.Println(slice)
	// slice[0] = "abay"
	// fmt.Println(slice)

	// slice := []int{3, 2, 1}
	// fmt.Println(slice)
	// sort.Ints(slice)
	// fmt.Println(slice)

	// array := []int{3, 2, 1}
	// fmt.Println(array)
	// sort.Ints(array)
	// fmt.Println(array)

	slice_str := []string{"madi", "adlet", "kayrat"}
	fmt.Println(slice_str)
	slice_str = append(slice_str, "amina")
	sort.Strings(slice_str)
	fmt.Println(slice_str)
}
