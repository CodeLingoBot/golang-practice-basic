package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main()  {
	dilbert.Salary -= 5000
	position := &dilbert.Position
	fmt.Println(position)
	*position = "Senior " + *position
	fmt.Println(*position)
}