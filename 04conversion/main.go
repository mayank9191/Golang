package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	// taking input from user
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)

	// doing operation on input(need to convert it first from string to float64 or int)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// 64 indicates that we need to convert it to float64 and TrimSpace removes the white spaces from the input because it is a string

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
