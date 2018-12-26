package main

import (
	"fmt"
)

var (
	counter int64
)
func main(){
	for i:=0;i<100 ;i++  {
		go func(){
			for i:=0; i< 10000; i++{
				counter++
			}
		}()
	}

	//var c string
	//fmt.Scanln(&c)
	fmt.Println(counter)
}

