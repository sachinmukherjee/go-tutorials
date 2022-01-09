package main

import "fmt"

func main() {

	//var mapName map[key_data_type]value_data_type
	var studentDetailsMap = make(map[int]string)
	studentDetailsMap[101] = "Sachin"
	studentDetailsMap[102] = "Sumeet"
	fmt.Println("Student Details Map ", studentDetailsMap)
	fmt.Println("Student Details for the Student 101", studentDetailsMap[101])
}
