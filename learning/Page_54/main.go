package main

import "fmt"

func main() {
	quantity := 4
	lenght, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(lenght*width, "square meters")
}
