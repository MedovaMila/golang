package main

import "fmt"

type Liters float64
type Gallons float64
type Millilitres float64

func (l Liters) ToMilliliters() Millilitres {
	return Millilitres(l * 1000)
}

func (m Millilitres) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", l, l.ToMilliliters())
	ml := Millilitres(500)
	fmt.Printf("%0.1f millilitres is %0.1f liters\n", ml, ml.ToLiters())
}
