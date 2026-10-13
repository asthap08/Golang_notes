package main

import "fmt"

func main() {
	//array of integers
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	//array of strings
	var fruits [3]string = [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits:", fruits)

	//accessing array elements
	fmt.Println("First number:", numbers[0])
	fmt.Println("Second fruit:", fruits[1])

	//modifying array elements
	numbers[2] = 10
	fmt.Println("Modified Numbers:", numbers)

	//array length
	fmt.Println("Length of numbers array:", len(numbers))
	fmt.Println("Length of fruits array:", len(fruits))

	//2d array
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Matrix:", matrix)

	//iterating over array elements
	fmt.Println("Iterating over numbers array:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println("Iterating over fruits array:")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

}
