package main

import (
	"fmt"
	"learning/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("/home/lmedova/go/bin/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
