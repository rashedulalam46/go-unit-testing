package main

import (
	"fmt"
)

func main() {

	cal:=  Calculator{}
	sum := cal.Add(5, 3)
	diff := cal.Subtract(5, 3)
	mul := cal.Multiply(5, 3)
	divided, err := cal.Divide(5, 3)
	if err != nil {
	 	fmt.Println("Error:", err)
	}
	

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Multification:", mul)
	fmt.Println("Devided:", divided)

}