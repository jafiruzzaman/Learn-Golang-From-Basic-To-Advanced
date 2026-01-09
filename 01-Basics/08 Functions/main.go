package main

import "fmt"
func main(){
	fmt.Println("========================================================== Basic Functions In Go-lang ==========================================================")
	greet()// Call greet function
	
	fmt.Println("========================================================== Parameterized Functions In Go-lang ==========================================================")
	var result int = add(10,20)
	fmt.Println("Result ",result)
}