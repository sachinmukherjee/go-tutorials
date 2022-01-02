package main

import "fmt"

func main() {
	printCutomerDetails("Sachin", "9009573764")
}

func printCutomerDetails(name string, contactNumber string) {
	fmt.Println("Customer Name ", name)
	fmt.Println("Customer Number ", contactNumber)
}
