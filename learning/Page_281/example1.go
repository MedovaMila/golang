package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

func nitroBoots(c *car) {
	c.topSpeed += 50
}
func main() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoots(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}