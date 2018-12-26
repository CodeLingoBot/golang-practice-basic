package main

import (
	"fmt"
	"sync"
)

var (
	countered int64
	mutex     sync.Mutex
)

func main() {
	for i := 0; i < 100; i++ {
		go func() { // Hàm vô danh
			for i := 0; i < 10000; i++ {
				mutex.Lock()
				countered++
				mutex.Unlock()
			}

		}()
	}

	fmt.Println(countered)
}
