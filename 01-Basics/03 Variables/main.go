package main

import "fmt"

func main() {
	fmt.Println("===================================== Variables in Go-lang =====================================")

	// What is a Variable?
	// A variable is a container to store values.
	// Analogy: If you want to store milk, you use a milk-bottle 🥛

	// 1. Standard variable declaration with explicit type
	var name string = "Go-lang"
	fmt.Println(name)

	// 2. Type inference (Go automatically detects the type)
	var goLang = "go-lang"
	fmt.Println(goLang)

	// 3. Short declaration (:=) — Go infers type automatically
	// ⚠️ Can only be used inside functions
	golang := "go-lang"
	fmt.Println(golang)
}
