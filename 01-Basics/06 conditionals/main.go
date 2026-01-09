package main

import "fmt"

func main()  {
	fmt.Println("======================================================= Conditionals In Go-lang =======================================================")
	// * 01. If-else Statement
	age:=1
	if age>=18 {
		fmt.Println("You are eligible for 🏎️ Drive ")
	}	else {
		fmt.Println("You Can't 🏎️  Drive")
	}
	age = 18
	voterId := true
	if age>=18 {
		if voterId {
			fmt.Println("You can Vote ")
		}	else{
			fmt.Println("You can't vote ")
		}
	} else{
			fmt.Println("You are not eligible.")
	}
	// * 02. Switch Case Statement
	var day string = "Friday"
	switch day {
	case "Saturday" :
		fmt.Println("The Day is ",day)
	case "Sunday":
		fmt.Println("The Day is ",day)
	case "Monday":
		fmt.Println("The Day is ",day)
	case "Tuesday":
		fmt.Println("The Day is ",day)
	case "Wednesday":
		fmt.Println("The Day is ",day)
	case "Thursday":
		fmt.Println("The Day is ",day)
	case "Friday":
		fmt.Println("The Day is ",day)
	default:
		fmt.Println("The Day is Not Valid")
	
	}
}