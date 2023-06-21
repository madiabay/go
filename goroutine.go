package main

import (
	"fmt"
)

// func main() {
// 	go say("888")
// 	go say("888")
// 	go say("888")
// 	fmt.Println("111")
// 	go say("888")
// 	go say("888")
// 	fmt.Println("333")
// 	fmt.Println("444")
// 	time.Sleep(2 * time.Second)
// }

// func say(say string) {
// 	fmt.Println(say)
// }
////////////////////////////////////////////////////////////////////////////////////////
// func main() {
// 	ch := make(chan int)

// 	go say(4, ch)
// 	fmt.Println(<-ch)

// }

// func say(a int, ch chan int) {
// 	fmt.Println(a)
// 	ch <- 7
// }

//////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	ch := make(chan int)
	go say(ch)

	for a := range ch {
		fmt.Println(a)
	}
}

func say(ch chan int) {
	for i := 5; i <= 10; i++ {
		ch <- i
	}

	close(ch)
}
