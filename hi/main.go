package main

import (
	"calc"
	"fmt"
	"greeting"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 3))
}
