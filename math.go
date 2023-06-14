package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// var a float32 = 9
	// var b float32 = 3
	// sum := a / b
	// fmt.Println(sum, reflect.TypeOf(sum))

	// var float_num float64
	// fmt.Scan(&float_num)
	// result := math.Sqrt(float_num)
	// fmt.Println(result, reflect.TypeOf(result))

	// var n float64 = 3.3195
	// var m float64 = 3.99
	// result := math.Ceil(n)
	// result2 := math.Floor(m)
	// fmt.Println(result, reflect.TypeOf(result))
	// fmt.Println(result2, reflect.TypeOf(result2))

	var round_num float64 = 5.49
	result := math.Round(round_num)
	fmt.Println(result, reflect.TypeOf(result))
}
