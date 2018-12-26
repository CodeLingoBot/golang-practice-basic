package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("init")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(3)

	go player("zen", court)
	go player("hiro", court)
	go player("aaa", court)

	court <- 1

	wg.Wait()
	fmt.Println("end game\n")

}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s win\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s failed\n", name)
			close(court)
			return
		}
		fmt.Printf("court - %d: %s\n", ball, name)
		ball++

		court <- ball
	}
}
