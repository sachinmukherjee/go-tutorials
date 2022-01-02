package main

import (
	"fmt"
)

func main() {
	var customerId int = 2500 //initializing at the time of declaration
	var validity int          //initializing later
	validity = 21
	var LogInMessage = "Welcome User" //initializing and declaring with type
	price := 100                      //shorthand declaration
	i, j, k := 20, 30, 40             //muliple variable
	_ = 55
	fmt.Println(customerId, validity, LogInMessage, price, i, j, k)
}
