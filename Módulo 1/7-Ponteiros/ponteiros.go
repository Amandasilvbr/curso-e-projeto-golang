package main

import "fmt"

func main() {
	num1 := 10
	num2 := &num1
	fmt.Println(num1, num2)

	num1++

	// * = desfaz a referência, pega o valor e não o lugar da memória 
	fmt.Println(num1, *num2)

}