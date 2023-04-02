package main

import "fmt"

func main(){
	arrayX := [5]int{} 
	arrayX= [5]int{1,2,3,4,5}
	fmt.Println(arrayX)

	sliceX := []int{}
	sliceX = []int{1,2,3,4}
	fmt.Println(sliceX)

	// Arrays internos, se passa da capacidade a capacidade é *2
	sliceY := make([]float32, 10, 15)
	fmt.Println(sliceY)
	fmt.Println(len(sliceY))
	fmt.Println(cap(sliceY))
}