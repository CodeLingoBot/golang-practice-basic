package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://demo6877814.mockable.io/a")
	if err != nil {
		log.Fatalf("err")
	}
	defer res.Body.Close()
	fmt.Println(res.Request)
}
