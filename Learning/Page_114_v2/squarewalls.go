package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//	for wall := 1; wall <=4; wall++ {
	fmt.Print("Enter widht and hight for first wall: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

//}
