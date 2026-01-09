package main

import "fmt"


func main(){
	fmt.Println("======================================================= Loops In Go-lang =======================================================")
	// what is loop?
	// loop is used to reduce repetitive works
	// ⚠️ There is only one single loop in Go-lang

	/* 1. Standard for loop
		for initialization;condition;updation {
				 * code
		}
	*/
	for i := 1; i <= 5; i++ {
		fmt.Print(i," ")
	}
	fmt.Println()

	// * 2. For loop as while loop
	/*
		* initialization --> var i int = 1
		* condition  --> for i<=5{
				* code
				* updation // 
		}
	*/

	fmt.Println("For Loop as while loop")
	i := 1
	for i<= 5 {
		fmt.Print(i," ")
		i++ // ⚠️ if-don't increment loop 
	}
	fmt.Println()
	fmt.Println("Loop with Break and Continue Statement")
	// * 3. loop with break and continue statement
	for i := 0; i <= 5; i++ {
		if i==2 {
			continue
		}
		if i==5 {
			break
		}
		fmt.Println("i = ",i)
	}
}