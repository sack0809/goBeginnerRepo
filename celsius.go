package main

import "fmt"

func main() {

	fmt.Println("Enter the Value:")

	var far float64
	fmt.Scanf("%f", &far)
	c := ((far - 32) * 5 / 9)
	fmt.Println(c)
}
