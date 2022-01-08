package main

import "fmt"

var name string
var mobileNo int64

func main() {
	fmt.Println("Name is ", getName())
	fmt.Println("Mobile no ", getMobileNo())
}

//function returning the value
func getName() string {
	name := "Sachin"
	return name
}

func getMobileNo() (mobileNo int64) {
	mobileNo = 9009573764
	return
}
