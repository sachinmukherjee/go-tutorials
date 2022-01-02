package main

import "fmt"

func main() {

	var plan string = "Prepaid"
	switch plan {
	case "Prepaid":
		fmt.Println("Prepaid Plan")
	case "Postpaid":
		fmt.Println("Postpaid Plan")
	default:
		fmt.Println("Undefined Plan!")
	}
}
