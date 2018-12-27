package main

import (
	"fmt"
)

func main(){
	ch:= make(chan int)

	go push(1, ch)
	go push(2, ch)
	go push(3, ch)
	go push(4, ch)
	go push(5, ch)
	go push(6, ch)
	go push(7, ch)
	go push(8, ch)
	go push(9, ch)

	fmt.Println(<-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch)
}

func push(num int, ch chan int){
	msg := num
	ch <- msg
}