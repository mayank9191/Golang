package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Hello, World"
	fmt.Println(welcome)

	// taking input from user

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our product: ")

	// comma ok || err ok syntax

	input, _ := reader.ReadString('\n') // "/n" symbolize till it read input

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input) // input is of string type to perform some actions on it we need to convert it first

}
