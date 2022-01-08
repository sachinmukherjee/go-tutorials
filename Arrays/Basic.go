package main

import "fmt"

func main() {
	//empty array
	var customerInvoice [10]int32
	//initilizing all the elements
	var studentName [3]string = [3]string{"Sachin", "Summet", "John"}
	//another way of initializing
	var name = [3]string{"Sachin", "Summet", "Jonny"}
	//multi dimensional array
	var multiDimension = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Customer Invoice ", customerInvoice)
	fmt.Println("Student Name ", studentName)
	fmt.Println("Name ", name)
	fmt.Println("MultiDimensional Array ", multiDimension)
}
