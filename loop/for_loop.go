package main

import"fmt"

func main(){
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//for loop as while loop
	j := 1
	for j <= 5 {
		fmt.Println("While Loop Iteration:", j)
		j++
	}
	//infinite loop
	for {
		fmt.Println("Infinite Loop Iteration")	
		break

	}
}
