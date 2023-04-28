package main

import "fmt"

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Sprinning")
}

type Coffepot string

func (c Coffepot) TurnOn() {
	fmt.Println("Powering up")
}

func (c Coffepot) Brew() {
	fmt.Println("Heating Up")
}

type Appliance interface {
	TurnOn()
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = Coffepot("LuxBrew")
	device.TurnOn()
}
