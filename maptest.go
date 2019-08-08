package main

import "fmt"

func main() {
	personSalary := map[string]int{
		"chaitanya": 11111,
		"aman":      22222,
	}
	personSalary["saurav"] = 3333
	fmt.Println("personsalaty map contents:", personSalary)
}
