package main

import "fmt"

func main() {

	var greeting = "Hello World"
	var ptrGreeting *string = &greeting //stores the address of the variable
	fmt.Println("Variable Address ", ptrGreeting)
}
