package main

import "fmt"

func main() {

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	delete(ages, "alice") // remove element ages["alice"]

	fmt.Println(ages["bob"])

	ages["bob"] = ages["bob"] + 1 // happy birthday!

	fmt.Println(ages)
}
