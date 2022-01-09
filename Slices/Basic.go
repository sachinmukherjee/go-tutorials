//slices are dynamic arrays

package main

import "fmt"

//TODO
func main() {
	customerContacts := make([]string, 7) //creating slice using make function. length is equal to size
	fmt.Println("Customer Contacts ", customerContacts)
	fmt.Println("Customer Contacts Length ", len(customerContacts))
	fmt.Println("Customer Contacts Capacity ", cap(customerContacts))

}
