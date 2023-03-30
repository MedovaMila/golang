package main

import "fmt"

func main() {
	ranks := map[string]int{"gold": 1, "siver": 2, "bronze": 3}
	for rank, medal := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", rank, medal)
	}
}
