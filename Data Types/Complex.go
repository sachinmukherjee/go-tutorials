package main

import "fmt"

func main() {
	var complex128 = 5 + 5i
	var complex64 = complex(2, 10)

	fmt.Println("Complex 128", complex128)
	fmt.Println("Real ", real(complex64), "Imaginary ", imag(complex64))
}
