package main

import "fmt"

func main(){
	fmt.Println("===================================== Data-Type In Go =====================================")
	// Integer
	var age int = 24
	fmt.Println("Age: ",age)

	// Float
	var price float64 = 66.98
	fmt.Println("Price: ",price)

	// string
	var name string = "Mohammad"
	fmt.Println("Name: ",name)

	// Boolean
	var isStudent bool = true
	fmt.Println("Student: ",isStudent)

	// Rune(Character)
	var grade rune = 'A'
	fmt.Println("Grade: ",string(grade))

	// Byte
	Byte := 255
	fmt.Println("Byte Value: ",Byte)
}