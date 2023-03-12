package main

import (
	"fmt"
)

func somar(num1, num2 int) int {
	return num1 + num2
}

func main(){
	soma := somar(10,20)
	fmt.Println(soma)
}
