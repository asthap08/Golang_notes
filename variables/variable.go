package main

import "fmt"

func main(){
	///////////
	var name string = "Vhits"
	fmt.Println("Hello " + name)

	///////////
	var names []string = []string{"Vhits", "John", "Doe"}
	fmt.Println(names)

	///////////
	name2 := "Vhits"
	fmt.Println("Hello " + name2)

	///////////
	var age = 25
	fmt.Println(age)

	///////////
	var isStudent bool = true
	if isStudent {
		fmt.Println("You are a student")
	} else {
		fmt.Println("You are not a student")
	}

	///////////
	var height float64 = 5.9
	fmt.Println(height)
}
