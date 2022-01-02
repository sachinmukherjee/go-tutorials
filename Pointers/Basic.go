package main

import "fmt"

func main() {

	var greeting = "Hello World"
	var ptrGreeting *string = &greeting
	fmt.Println("Variable Address ", ptrGreeting)
}
