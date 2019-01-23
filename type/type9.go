package main

import (
	"encoding/json"
	"fmt"
)

var Config config

type config struct {
	Host string `json:"host"`
	Name string `json:"name"`
}
type response2 struct {
	Page   int      `json:"pagesss"`
	Fruits []string `json:"fruit"`
}

func main() {
	str := `{"pagesss": 1, "fruit": ["apple", "peach"]}`
	str1 := `{"host": "123", "name": "1233123"}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//l:= new(config)
	json.Unmarshal([]byte(str1), &Config)
	fmt.Println(Config)
}
