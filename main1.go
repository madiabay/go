package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World")
	// var age uint8 = 255
	// fmt.Println(age)
	// var float_number float64 = 34.555
	// fmt.Println(float_number)

	// age := 16.1
	// var age2 float32 = 16.1
	// var age int8 = 16
	// fmt.Println(reflect.TypeOf(age), reflect.TypeOf(age2))

	// var name string = "Madi"
	// fmt.Println(name, reflect.TypeOf(name))
	// name2 := "Madi"
	// fmt.Println(name2, reflect.TypeOf(name2), len(name2))

	// var surname string
	// fmt.Println("What is your surname?")
	// fmt.Scan(&surname)
	// fmt.Println("Hello " + surname + "!!!")
	// var age uint8
	// fmt.Println("How old are you?")
	// fmt.Scan(&age)
	// fmt.Println(surname + " is " + fmt.Sprint(age) + " years old!")

	var a int8 = 2
	var b float32 = float32(a)
	fmt.Println(b, reflect.TypeOf(b))
}
