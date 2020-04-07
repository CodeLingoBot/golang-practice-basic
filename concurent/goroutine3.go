package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorker = 4
	numTask   = 10
)

var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().Unix())
}

func main(){
	tasks := make(chan string, numTask)

	wg.Add(numWorker)
	for gr:=1;gr<= numWorker ;gr ++  {
		go worker(tasks, gr)
	}

	for t:=1;t<=numTask ;t++  {
		tasks <- fmt.Sprintf("job at - %d", t)
	}

	close(tasks)
	wg.Wait()
}

func worker(task <-chan string, worker int){
	defer wg.Done()

	for{
		task, ok:= <- task
		if !ok{
			fmt.Printf("1 worker - %s - done\n", task)
			return
		}

		fmt.Printf("2 worker %d start %s\n", worker, task)

		sleep:= rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("3 worker %d: done %s\n", worker, task)

	}
}