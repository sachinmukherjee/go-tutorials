package main

import "fmt"

func main() {
	var plan string = "Postpaid"

	if plan == "Prepaid" {
		fmt.Println("Existing plan is prepaid")
	} else if plan == "Postpaid" {
		fmt.Println("Existing plan is postpaid")
	} else {
		fmt.Println("Plan is undefined")
	}
}
