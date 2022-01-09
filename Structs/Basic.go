package main

import "fmt"

type Customer struct {
	name        string
	age         int
	phoneNumber int64
}

func main() {
	//initializing technique 1
	var customer1 = Customer{"Sachin", 27, 9009573764}
	//initializing technique 2
	var customer2 = Customer{name: "Sumeet", age: 30}
	fmt.Println("Customer 1", customer1)
	fmt.Println("Customer 2", customer2)
}
