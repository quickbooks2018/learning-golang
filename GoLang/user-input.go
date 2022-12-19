package main

import "fmt"

func main() {
	var fName string
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&fName)

	var lName string
	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lName)

	fmt.Print("Full Name is: " + fName + " " + lName)
}
