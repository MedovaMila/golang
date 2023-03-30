package main

import (
	"fmt"
	datafile "learning/datafile/new"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, grade := range counts {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
