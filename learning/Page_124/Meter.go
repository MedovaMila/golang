package main

import "fmt"

var meterPerLiter float64

func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / meterPerLiter
}

func main() {
	meterPerLiter = 10.0
	fmt.Printf("%.2f\n", paintNeeded(4.2, 3.0))
}
