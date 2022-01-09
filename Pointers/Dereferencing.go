package main

import "fmt"

func main() {
	var greeting string = "Hello World"
	var ptrGreeting = &greeting //stores the address
	fmt.Println("Value of greeting ", greeting)
	fmt.Println("Address of greeting ", &greeting)
	fmt.Println("value of ptr greeting ", ptrGreeting)
	fmt.Println("accessing value of greeting using pointer ptr greeting ", *ptrGreeting) //fetch the value pointer refers to
}
