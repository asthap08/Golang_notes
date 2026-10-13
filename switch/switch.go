package main

import "fmt"

func main() {
	//switch statement
	var day int = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	//multiple condition switch statement
	var grade string = "B"
	switch grade {
	case "A", "B", "C":
		fmt.Println("Pass")
	case "D", "F":
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid grade")
	}

	//type switch
	var value interface{} = 42
	switch v:=value.(type){
	case int:
		fmt.Println("integer value is",v)
	case string:
		fmt.Println("string value is",v)
	default:
		fmt.Println("unknown type")
	}
}
