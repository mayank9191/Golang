package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	//you cannot wrie the function inside the function

	fmt.Println("Result is:", result)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proadder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namstey from golang")
}
