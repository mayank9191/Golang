package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	// Lower limit is inclusive and Highher is exclusive in siles(append)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 768
	highScore[3] = 900

	fmt.Println(highScore)

	highScore = append(highScore, 555, 666, 321)

	//  sorting of slices

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
