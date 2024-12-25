package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Potato"
	fruitList[3] = "Mango"
	fmt.Println("List of fruits: ", fruitList)
	fmt.Println("List of fruits: ", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}

	fmt.Println{"List of vegetables:", vegList}
	fmt.Println{"List of vegetables:", len(vegList)}
}
