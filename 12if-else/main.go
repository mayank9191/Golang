package main

import "fmt"

func main() {

	fmt.Println("Welcome to if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		fmt.Println("Regular user")

	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

}
