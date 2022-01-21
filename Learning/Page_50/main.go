package main

import "fmt"

func main() {
	var quantity int
	var lenght, widht float64
	var custorName string

	quantity = 4
	lenght, widht = 1.2, 2.4
	custorName = "Daimon Cole"

	fmt.Println("custorNam")
	fmt.Println(custorName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println(lenght*widht, "square meters")
}
