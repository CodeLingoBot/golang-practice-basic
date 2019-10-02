package main

import (
	"fmt"
	"time"
)

func main() {
	thoaArr := [5]string{"nguyen", "thi", "kim", "thoa"}

	thoaChan := make(chan string)

	go kim(thoaArr, thoaChan)

	go thoa(thoaChan)

	<-time.After(time.Second * 1)
}

func kim(th [5]string, thoachan chan<- string) {
	for _, item := range th {
		thoachan <- item
	}
}

func thoa(thoachan <-chan string){
	for i:=0; i< 4; i++{
		f := <- thoachan
		fmt.Println(" "+ f)
	}
}
