package main

import "fmt"

type User struct {
	username string
	email    string
	age      int32
}

func change(u *User) {
	u.username = "Kayroshhhhhhhh"
}

func main() {
	// var user User = User{age: 21, username: "Madi", email: "abaymadi1@gmail.com"}
	// fmt.Println(user)
	// fmt.Println(user.age, reflect.TypeOf(user.age))

	// user2 := User{"Adlet", "adlet@bk.ru", 25}
	// fmt.Println(user2, user2.username)

	// user3 := User{username: "John", age: 22, email: "john@gmail.com"}
	// fmt.Println(user3)
	// user3.username = "john_doe"
	// user3.age = 35
	// fmt.Println(user3, reflect.TypeOf(user3))

	user4 := User{"Kayrat", "kayrat01@bk.ru", 10}
	fmt.Println(user4)
	change(&user4)
	fmt.Println(user4)
}
