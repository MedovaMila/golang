package main

import "fmt"

type Whistle string

func (w Whistle) MakeSoound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSoound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSoound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSoound()
}

func play(n NoiseMaker) {
	n.MakeSoound()

}

func main() {
	play(Robot("Botco Ambler"))
}
