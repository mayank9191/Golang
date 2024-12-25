package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to switch case in golang")

	rand.NewSource(time.Now().UnixNano())

	diceNumber := rand.Intn(5) + 1

	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("you can move 2 spot")

	case 3:
		fmt.Println("you can move 3 spot")

	case 4:
		fmt.Println("you can move 4 spot")

	case 5:
		fmt.Println("you can move 5 spot")

	case 6:
		fmt.Println("you can move 6 spot and roll the dice again")

	default:
		fmt.Println("What was that!")
	}
}
