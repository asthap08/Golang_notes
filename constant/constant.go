package main

import "fmt"

func main(){
	const name string = "Vhits"
	fmt.Println("Hello " + name)

	const age int = 25
	fmt.Println(age)

	const isStudent bool = true
	if isStudent {
		fmt.Println("You are a student")
	} else {
		fmt.Println("You are not a student")
	}

	const height float64 = 5.9
	fmt.Println(height)
}
