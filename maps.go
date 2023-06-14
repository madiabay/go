package main

import (
	"fmt"
)

func main() {
	// var currency map[string]float32 = map[string]float32{
	// 	"rub":    5.5,
	// 	"tenge":  1,
	// 	"dollar": 475.8,
	// 	"euro":   500.5,
	// }
	// fmt.Println(currency)
	// fmt.Println(currency["dollar"])

	// var fio map[string]string = map[string]string{
	// 	"surname": "Abay",
	// 	"name":    "Madi",
	// }
	// fi2 := map[string]string{
	// 	"surnam2": "Aba2",
	// 	"nam2":    "Mad2",
	// }
	// fmt.Println(fio, reflect.TypeOf(fio))
	// fmt.Println(fi2, reflect.TypeOf(fi2))

	// alph := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }
	// fmt.Println(alph)
	// delete(alph, "a")
	// fmt.Println(alph)
	// delete(alph, "z")
	// fmt.Println(alph)

	dict_map := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	fmt.Println(dict_map)
	element, status := dict_map[3]
	fmt.Printf("Element-%d, Status-%t\n", element, status)
	element2, status2 := dict_map[4]
	fmt.Printf("Element2-%d, Status2-%t\n", element2, status2)
	// fmt.Println(reflect.TypeOf(element2), reflect.TypeOf(status2))
}
