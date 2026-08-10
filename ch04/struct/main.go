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

func EmployeeByID(id int) *Employee {
	return &Employee{}
}

func EmployeeByID1(id int) Employee {
	return Employee{}
}

var dilbert Employee

func main() {

	id := 1
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id1 := EmployeeByID1(id)
	id1.Salary = 1
}
