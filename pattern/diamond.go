package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if j==3 || i == 1 || i == 6 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
