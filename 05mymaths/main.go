package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	// fmt.Println("The sum is:", mynumberOne+int(mynumberTwo))

	fmt.Println("Choose one of the follwing method to generate random number")
	fmt.Println("1> Random number generator using math algo")
	fmt.Println("2> Random number generator using crypto algo")

	reader := bufio.Newreader(os.Stdin)

	choice, _ := reader.ReadString('\n')

	if choice == "1\n" {

		// random number generator using math algo

		rand.NewSource(time.Now().UnixNano())

		fmt.Println(rand.Intn(5) + 1)

		fmt.Println(time.Now().UnixNano())
	}

	if choice == "2\n" {
		// random number generator using crypto algo
		myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))

		fmt.Println(myRandomNum)

	}

}
