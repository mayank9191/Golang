package main

import "fmt"

const LoginToken string = "sdfsd" //capital L in LoginToken symbolize that it is (public) and can be anywhere in file

func main() {

	//  Declaration and Initialization of variables

	// STRING

	var username string = "Mayank"

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//  BOOLEAN

	var isLoggedIn bool = false

	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// INTEGER

	var smallVal uint8 = 255

	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// FLOAT

	var smallFloat float64 = 255.886877665676

	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default value of int is 0 and some aliases

	var someVariables int

	fmt.Println(someVariables)
	fmt.Printf("Variable is of type: %T \n", someVariables)

	//default value of string is " "
	var Variable string

	fmt.Println(Variable)
	fmt.Printf("Variable is of type: %T \n", Variable)

	// implicit type(without specifying type of the variable)

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style(not applicable for global variables)

	numberOfUser := 300000 //walrus syntax
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
