package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data_byte := []byte("Madi Adlet Kayrat Amina")
	e := ioutil.WriteFile("text.txt", data_byte, 0600)

	if e != nil {
		fmt.Println("I can not create this file\n", e)
	}

	file_data, err := ioutil.ReadFile("text.txt")

	if err != nil {
		fmt.Println("I can not read this file\n", err)
	}

	fmt.Println(string(file_data))

	os.Remove("text.txt") // для удаления файла
}
