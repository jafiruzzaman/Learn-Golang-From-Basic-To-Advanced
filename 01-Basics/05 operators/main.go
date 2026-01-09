package main

import "fmt"

func main()  {
	a:= 10
	b:= 3
	// * 01. Arithmetic Operator
	fmt.Println("======================================================= Arithmetic Operator =======================================================")
	fmt.Println("Add: ",(a+b))
	fmt.Println("Subtract: ",(a-b))
	fmt.Println("Multiplication: ",(a*b))
	fmt.Println("Division: ",(a/b))
	fmt.Println("Mod: ",(a%b))

	// * 02. Assignment Operator
	fmt.Println("======================================================= Assignment Operator =======================================================")
	a+=5
	fmt.Println("Add and Assign",a)
	a-=5
	fmt.Println("Subtract and Assign",a)
	a*=5
	fmt.Println("Multiply and Assign",a)
	a/=5
	fmt.Println("Divide and Assign",a)
	a%=5
	fmt.Println("Modulo and Assign",a)



	// * 03. Comparison Operator
	fmt.Println("======================================================= Comparison Operator =======================================================")
	fmt.Println("a > b:", a>b)
	fmt.Println("a >= b:", a>=b)
	fmt.Println("a < b:", a<b)
	fmt.Println("a <= b:", a<=b)
	fmt.Println("a == b:", a==b)
	fmt.Println("a != b:", a!=b)

	// * 04 Logical Operator
	fmt.Println("======================================================= Logical Operator =======================================================")
	x := true
	y := false
	fmt.Println(x && y) // false
	fmt.Println(x || y) // true
	fmt.Println(!x)     // false
}