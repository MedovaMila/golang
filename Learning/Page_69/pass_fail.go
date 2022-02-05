package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Println(input)

	//if grade == 100 {
		fmt.Println("Perfect!")
	} else if grade >= 60 {
		fmt.Println("You pass.")
	} else {
		fmt.Println("You fail!")
	}
}
