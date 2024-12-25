package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	// no inheritance in golang

	// no class in golang

	mayank := User{"Mayank", true, 20, "kumar@go.dev"}
	fmt.Println(mayank)

	fmt.Printf("Name is %v\n", mayank.Name)

}

type User struct {
	Name   string
	Status bool
	Age    int
	Email  string
}
