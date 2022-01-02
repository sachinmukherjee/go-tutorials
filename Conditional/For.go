package main

import "fmt"

func main() {
	for count := 1; count <= 10; count++ {
		fmt.Println("Value of count is ", count)
	}

	for _, element := range "Sachin" {
		fmt.Println(string(element))
	}
}
